[Loading CUE](https://cuetorials.com/go-api/loading/)

Loading CUE aims to mimic the process the CLI uses to construct a value from entrypoints or arguments. This section and subsections go over the loading process and configuration. The next section deals with managing the the gaps and differences to fully replicate the CLI behavior.

The Loading Process
The CUE loader or cue/load package can be used to load CUE much like the CLI. The process consists of:

specifying the entrypoints or arguments
setting up the configuration
calling load.Instances
turning or building the Instances into a CUE Value
We will first walk through a basic example of this process then look at details and advanced usage.


Configuration
Overlay
Data Files
