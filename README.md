# argument-analysis-research

A collection of code used for the research project at ARG-tech aiming for argument analysis and natural language processing

## About

The main work for this research project will happen inside [Colaboratory](https://colab.research.google.com/drive/1iGL_J01I-SAtw2HG8uoJMLgYhYqMzzAK) but due to lack of existing tooling for what we are trying to achieve here we also have to try non Python solutions which should reside in one collective place (so this repo).

## Segmenting

Many of the answers in our data set for this project are long and contain multiple statements. To correctly analyze them, we need to segment those.

As a first experimental solution there is the simple cli tool inside [cmd/segmenter/segmenter.go](cmd/segmenter/segmenter.go) which is using [github.com/jdkato/prose](https://github.com/jdkato/prose) for doing the first steps of segmenting.

To run it simply provide your string as first argument:

```bash
go run cmd/segmenter/*.go "Hello darkness my old friend"
```

## Argument Analysis API

Some parts of this repo get combined into a simple JSON API for use in other places (like Colaboratory).
This currently includes segmenting and keyword extraction.

To analyze a piece of text send it to the API Server (dev running at https://research.democracy.ovh/analyze) as a POST in the format of: 

```json
{
  "input": "Hello darkness my old friend"
}
```

This will apply the segmenting mentioned above and then extract keywords by applying RAKE.

The response will look like this:

```json
{
  "content": "Hello darkness my old friend",
  "segments": [
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

Find the corresponding license attached in [LICENSE](LICENSE).
This project is meant for experimenting which means using and vendoring in foreign code where suitable.
This license does only apply to code written by contributors, external libraries might differ.
