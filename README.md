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

For a detailed API documentation, please refer to the APIBlueprint located in [apiary.apib](apiary.apib) or check out the interactive docs at [Apiary](https://argumentanalysisresearch.docs.apiary.io/#).

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
