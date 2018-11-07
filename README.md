# argument-analysis-research

A collection of code used for the research project at ARG-tech aiming for argument analysis and natural language processing

## About

The purpose of this project is to provide tools and services used in our argument analysis research.

Parts of this research project will happen inside [Colaboratory](https://colab.research.google.com/drive/1iGL_J01I-SAtw2HG8uoJMLgYhYqMzzAK) as well. So this repository is a dependency there and should be treated accordingly.

## Segmenting

Many of the answers in our data set for this project are long and contain multiple statements. To correctly analyze them, we need to segment those.

As a first experimental solution there is the simple cli tool inside [cmd/cli/segmenter.go](cmd/cli/segmenter.go) which is using [github.com/jdkato/prose](https://github.com/jdkato/prose) for doing the first steps of segmenting.

To run it simply provide your string as first argument:

```bash
go run cmd/segmenter/*.go "Hello darkness my old friend"
```

## Argument Analysis API

Some parts of this repo get combined into a simple JSON API for use in other places (like Colaboratory).
This currently includes segmenting, keyword extraction as well as ADW based segment comparison.

### Segmenting Endpoint

Segmenting currently uses [github.com/jdkato/prose](https://github.com/jdkato/prose) until our research leads to more sophisticated methods.

#### Request

To segment a piece of text send it to the API Server (dev running at https://research.democracy.ovh/argument/segment) as a POST in the format of: 

```json

{
  "input": "Hello darkness my old friend"
}

```

#### Response

```json
{
  "content": "Hello darkness my old friend",
  "segments": [
    {
      "text": "Hello darkness my old friend"
    }
  ]
}

```

### Keyword Extraction Endpoint

For Keyword extraction we currently use RAKE. The endpoint adds found keywords to the passed in segment.

#### Request

To get keywords for a segment, send it to the API Endpoint (dev running at https://research.democracy.ovh/argument/keyword) as POST in the format of:

```json

{
  "content": "Hello darkness my old friend",
  "segments": [
    {
      "text": "Hello darkness my old friend"
    }
  ]
}

```

#### Response

```json

{
  "text": "Hello darkness my old friend",
  "keywords": [
    {
      "key": "old friend",
      "value": 4
    },
    {
      "key": "darkness",
      "value": 1
    }
  ]  
}

```

### ADW API

This almost primitive service wraps around [ADW](https://github.com/pilehvar/ADW) to expose a very simple API for checking text similarity of a json request.

The ADW API is written in Kotlin and can be found [here](src/main/kotlin/com/github/cdl/adw/Service.kt).
It has certain setup requirements:

* [config/](config/): Contains the ADW and related configs which are required at runtime.
* [resources/](resources/): Contains various directories that are required at runtime. All except the signatures are included in this repo.
  * [resources/signatures](resources/signatures): Should contain the exported signature file which can be downloaded from the original [link](http://lcl.uniroma1.it/adw/ppvs.30g.5k.tar.bz2) in the ADW repo.

#### Request

```json
{
  "text1": "We should stop eating meat",
  "text2": "Don't buy so much stuff"
}
```

#### Response

```json
{
  "result": 0.507
}
```

### Linker API

This service takes a list of segmented Documents and returns a list/matrix of all links between segments and indirectly their parent answers as well.
The rating API to use is flexible and can be set by adding the rating API Url in the request. The Rating API has to match the interface of the ADW API above.

For now the Linking process is as primitive as it could be and just compares **all** segments against each other.
This is naturally shitty inefficient and needs a lot more love. Also the default ADW API is very slow which makes the request even slower.

Immediate Todos are:

* Add async call so the request does not need to wait
* Use a parallel approach for comparing all segments instead of doing all in a row
* Preprocess and clean data to speed up the process
* Speed, speed, speed
* Add two way comparison as some stuff like the ADW API rate differently depending on option order

**Note: This API ist highly experimental and just one experiment of many others. It might change at any time and does neither meet any of our code and quality standards.**

#### Request

The request has to contain the rating api url as well as a list of **segmented** documents.
The order of the documents will be retained and represent the order of the response data.

```json
{
  "rater": "https://research.democracy.ovh/argument/linker",
  "documents":[
    {
      "content": "Government incentives to upgrade coal fired power plants and coal processing plants to make them more efficient and less polluting  Replace old nuclear power plants with newer, safer more efficient designsLarger p ublic investment in wind, solar, and where feasible geothermal energy Upgrade the power grid to be more efficient  (because a large part of our carbon footprint is electricity generation) Fuel efficiency standards for large trucks and not just passenger cars",
      "segments": [
        {
          "text": "Government incentives to upgrade coal fired power plants and coal processing plants to make them more efficient and less polluting  Replace old nuclear power plants with newer, safer more efficient designsLarger p ublic investment in wind, solar, and where feasible geothermal energy Upgrade the power grid to be more efficient  (because a large part of our carbon footprint is electricity generation) Fuel efficiency standards for large trucks and not just passenger cars"
        }
      ]
    },
    {
      "content": "Keep encouraging regulations reducing pollutants into our atmosphere.",
      "segments": [
        {
          "text": "Keep encouraging regulations reducing pollutants into our atmosphere."
        }
      ]
    },
    {
      "content": "The eruption of two supervolcanoes would take care of it. At least for a while.",
      "segments": [
        {
          "text": "The eruption of two supervolcanoes would take care of it."
        },
        {
          "text": "At least for a while."
        }
      ]
    },
    {
      "content": "If every human being on the planet dropped dead right now, the worst effects of climate change would still occur. There's no stopping it at this point. Anyone who tries to tell you differently is deluded.",
      "segments": [
        {
          "text": "If every human being on the planet dropped dead right now, the worst effects of climate change would still occur."
        },
        {
          "text": "There's no stopping it at this point."
        },
        {
          "text": "Anyone who tries to tell you differently is deluded."
        }
      ]
    }
  ]
}
```

#### Response

The response (for now) just returns a matrix of related answers based on their segments. To be more precise it just contains the links between segments based on their document.
The Links array contains the document links in the order they were passed to the API. The document then contains an array of segment links with each segment containing various links directly addressing the related segments.

In this sample data you can see, that the data passed in is in no way related to each other as all segments just link to themselves (with a value of 1). My fault, bad test data.  

```json
{
    "Links": [
        [
            [
                {
                    "Document": 0,
                    "Segment": 0,
                    "Dist": 1
                }
            ]
        ],
        [
            [
                {
                    "Document": 1,
                    "Segment": 0,
                    "Dist": 1
                }
            ]
        ],
        [
            [
                {
                    "Document": 2,
                    "Segment": 0,
                    "Dist": 1
                }
            ],
            []
        ],
        [
            [
                {
                    "Document": 3,
                    "Segment": 0,
                    "Dist": 1
                }
            ],
            [
                {
                    "Document": 3,
                    "Segment": 1,
                    "Dist": 1
                }
            ],
            [
                {
                    "Document": 3,
                    "Segment": 2,
                    "Dist": 1
                }
            ]
        ]
    ]
}
```

## Contributing

This repository is part of a 2 week research project.
If you are part of the [CDL](https://canonicaldebatelab.com/) or [DPT](http://digitalpeacetalks.com/) feel free to request access and contribute.
Sync is happening through Mail or our [Slack](https://join.slack.com/t/canonicaldebatelab/shared_invite/enQtMzEzOTU3NzYyMDY3LTI4YzUxM2I0MjFjZDNlMzQxZDM4YTgwNDNlMTY3YWQwNjJhYjk0ODE1MGU5NzQ2MTAyNTFhZWRhMGNjMjAxNmE).

If you are an external party and interested in contributing, feel free to get in touch.
While this repo is just for the research project, our effort does not stop after two weeks and we are happy to have people join the mission.

## License

### General

Find the corresponding license attached in [LICENSE](LICENSE).
This project is meant for experimenting which means using and vendoring in foreign code where suitable.
This license does only applies to code written by contributors, external libraries can differ and got their licenses stored in the respective vendor directory.

### Code for ADW

The licensing from the original project applies without any changes.

>ADW (Align, Disambiguate and Walk) -- A Unified Approach for Measuring Semantic Similarity.
>
>Copyright (c) 2014 Sapienza University of Rome.
>All Rights Reserved.
>
>This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY;
>without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
>
>If you use this system, please cite the following paper:
>
>> M. T. Pilehvar, D. Jurgens and R. Navigli. Align, Disambiguate and Walk: A Unified Approach for Measuring Semantic Similarity.
>> Proceedings of the 51st Annual Meeting of the Association for Computational Linguistics (ACL 2013), Sofia, Bulgaria, August 4-9, 2013, pp. 1341-1351.

### Code for wrapping around ADW

The code for wrapping around the ADW implementation is licensed through the same [LICENSE](LICENSE) as the original library.

Thanks to the respective authors and developers for providing their work.
