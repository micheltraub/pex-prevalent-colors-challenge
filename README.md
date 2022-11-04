# PEX: Prevalent Colors Challenge
[![Go](https://github.com/micheltraub/pex-prevalent-colors-challenge/actions/workflows/go.yml/badge.svg)](https://github.com/micheltraub/pex-prevalent-colors-challenge/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/micheltraub/pex-prevalent-colors-challenge/branch/main/graph/badge.svg?token=TDYEJLZMR2)](https://codecov.io/gh/micheltraub/pex-prevalent-colors-challenge)

---
- [PEX: Prevalent Colors Challenge](#pex-prevalent-colors-challenge)
  - [Running the application](#running-the-application)
  - [Challenge description](#challenge-description)
    - [The first part of the challenge description says:](#the-first-part-of-the-challenge-description-says)
      - [Proposed solution](#proposed-solution)
    - [The second part of the challenge description says:](#the-second-part-of-the-challenge-description-says)
      - [Proposed solution](#proposed-solution-1)
  - [Results](#results)

---

## Running the application
This application can be launched with **Go**, run:
```shell
go mod download 
go run cmd/app/main.go
```

---
## Challenge description

### The first part of the challenge description says:
>Below is a list of links leading to an image. Read this list of images and find the 3 most prevalent colors in the RGB scheme in >hexadecimal format (#000000 - #FFFFFF) for each image, then write the result into a CSV file in the form of url,color,color,color.

#### Proposed solution
To configure the input location, the input filename, the output location, and the output filename you can change the `.Env` file modifying the following default values:

```dosini
#INPUT Location
INPUT_PATH=./test/data/
INPUT_FILENAME=input1000.txt

#OUTPUT location
OUTPUT_PATH=./out/
CSV_OUTPUT_FILENAME=colors_output.csv
```

To calculate the 3 most prevalent colors the system provides an *interface* [`prevalentcolor`] because we can get that result using different approaches.
The following approches are:
- The **accurate** one, that counts the HEX code of all colors and shows the 3 that repeats the most.
- The **average** one, that uses `KMeans clustering` to calculate the prevalent colors (TODO: not implemented yet)

To change the approach you can set in the `.Env` file changing the following:
```dosini
#Calculating prevalent colors mode: ACCURATE or AVERAGE
PREVALENT_MODE=ACCURATE
```
To better perform, specially with big images, we can reduce the images before calculating the prevalent colors. Doing that makes it faster but less accurate, even that we use interpolation to keep the ratio. To set the config to reduce the image, in the `.Env` file:
```dosini
#Reducing image improve performance but lose accuracy
REDUCE_IMAGES=true
```
### The second part of the challenge description says:
>Please focus on speed and resources. The solution should be able to handle input files with more than a billion URLs using limited >resources (e.g., 1 CPU, 512MB RAM). Keep in mind that there is no limit on the execution time, but make sure you are utilizing the >provided resources as much as possible at any time during the program execution. 

#### Proposed solution
The application is focused on speed. So the pipeline used allows for every *URL* in the input file, the application runs a **goroutine** to fetch and calculate the 3 most prevalent colors of that image. Those are running in parallel and sending the results  The result is sent to the `csvLine channel`. 

In parallel, the application runs **goroutines** to read the `csvLine channel` and append to the output CSV file

This approach is the fastes but will use all the resources available. Below some **CPU** and **memory** usage.

---

## Results

To enable **CPU** and **memory** monitoring, set the following variables in the `.Env` file:
```dosini
#MONITORING for performance
ENABLE_CPU_MONITOR=true
ENABLE_MEMORY_MONITOR=true
CPU_PPROF_FILENAME=cpu.pprof
```

Processing `test/data/input1000.txt`: 
- With `REDUCE_IMAGES` set to **true**:
  - Memory usage: 
  ```dosini
    2022/11/03 23:32:33 Process took 2m0.7072298s
    2022/11/03 23:32:33 
    2022/11/03 23:32:33 Alloc: 4088 MB, TotalAlloc: 16461 MB, Sys: 10396 MB
    2022/11/03 23:32:33 Mallocs: 192031115, Frees: 148867029
    2022/11/03 23:32:33 HeapAlloc: 4088 MB, HeapSys: 10019 MB, HeapIdle: 5924 MB
    2022/11/03 23:32:33 HeapObjects: 43164086
    2022/11/03 23:32:33
    ```
  - CPU PPROF top10 (run: `go tool pprof cpu.pprof` and after that `top10`):
  
  ```shell
  Type: cpu
    Time: Nov 3, 2022 at 11:30pm (EDT)
    Duration: 120.84s, Total samples = 1.91hrs (5695.12%)
    Entering interactive mode (type "help" for commands, "o" for options)
    (pprof) top10
    Showing nodes accounting for 6754.48s, 98.15% of 6881.71s total
    Dropped 889 nodes (cum <= 34.41s)
    Showing top 10 nodes out of 19
        flat  flat%   sum%        cum   cum%
        flat  flat%   sum%        cum   cum%
    6714.89s 97.58% 97.58%   6714.94s 97.58%  runtime.cgocall
        34.84s  0.51% 98.08%     38.13s  0.55%  github.com/nfnt/resize.resizeYCbCr
            4s 0.058% 98.14%     34.92s  0.51%  image/jpeg.(*decoder).processSOS
        0.71s  0.01% 98.15%     49.24s  0.72%  pex-prevalent-colors-challenge/internal/app/accurateprevalent.(*AccuratePrevalentColor).CalculatePrevalentColors
        0.01s 0.00015% 98.15%     37.31s  0.54%  image/jpeg.(*decoder).decode
        0.01s 0.00015% 98.15%   6685.10s 97.14%  main.PersistToCsvFile.func1
        0.01s 0.00015% 98.15%   6685.08s 97.14%  pex-prevalent-colors-challenge/internal/app/csv.AppendToCsvFile
        0.01s 0.00015% 98.15%   6714.89s 97.58%  syscall.SyscallN
            0     0% 98.15%     55.85s  0.81%  image.Decode
            0     0% 98.15%     37.34s  0.54%  image/jpeg.Decode
    ```
- With `REDUCE_IMAGES` set to **false**:
    - Memory usage: 
    ```dosini
    022/11/03 23:46:28 Process took 5m59.3611309s
    2022/11/03 23:46:28 
    2022/11/03 23:46:28 Alloc: 2958 MB, TotalAlloc: 50174 MB, Sys: 12046 MB
    2022/11/03 23:46:28 Mallocs: 4776952917, Frees: 4734262603
    2022/11/03 23:46:28 HeapAlloc: 2958 MB, HeapSys: 11457 MB, HeapIdle: 8022 MB
    2022/11/03 23:46:28 HeapObjects: 42690314
    2022/11/03 23:46:28
    ```
    - CPU PPROF top10 (run: `go tool pprof cpu.pprof` and after that `top10`):
  
    ```shell
    Type: cpu
    Time: Nov 3, 2022 at 11:40pm (EDT)
    Duration: 359.53s, Total samples = 1457.22s (405.31%)
    Entering interactive mode (type "help" for commands, "o" for options)
    (pprof) top10
    Showing nodes accounting for 961.33s, 65.97% of 1457.22s total
    Dropped 835 nodes (cum <= 7.29s)
    Showing top 10 nodes out of 86
        flat  flat%   sum%        cum   cum%
    207.32s 14.23% 14.23%    207.35s 14.23%  runtime.cgocall
    164.17s 11.27% 25.49%    333.20s 22.87%  runtime.mapaccess1_faststr
    118.71s  8.15% 33.64%    412.09s 28.28%  fmt.(*pp).doPrintf
    107.74s  7.39% 41.03%    113.47s  7.79%  image.(*YCbCr).YCbCrAt
        95.29s  6.54% 47.57%     95.29s  6.54%  aeshashbody
        81.29s  5.58% 53.15%    168.85s 11.59%  fmt.(*fmt).fmtInteger
        61.60s  4.23% 57.38%     61.60s  4.23%  memeqbody
        43.36s  2.98% 60.35%   1176.70s 80.75%  pex-prevalent-colors-challenge/internal/app/accurateprevalent.(*AccuratePrevalentColor).CalculatePrevalentColors
        41.65s  2.86% 63.21%    235.90s 16.19%  fmt.(*pp).printArg
        40.20s  2.76% 65.97%     41.19s  2.83%  image.(*NRGBA).NRGBAAt
    ```
---

## Improvements
As I am new in GO I had to learn many things, so I didnt had the time to do all that I wanted to do. So my goal was to apply some concepts that I know but are important but there are many things that are incomplete or that could be improved. Here is a list of things that I would do:

- Implement the average prevalent color calcuation using KMeans Clustering
- Limit the number of goroutines using things like runtime.NumCPU() and concurrency best practices (https://go.dev/blog/pipelines)
- Improve Error handler: create error types
- Add log level management
- Improve test coverage
- Reorganize my code following go project structure and som refactor using best practices (double checking naming conventions, dependecy injections, apply some patterns, etc...). Some references:
    - https://github.com/golang-standards/project-layout
    - https://golangbyexample.com/all-design-patterns-golang/
    - https://go.dev/blog/pipelines
    - https://go.dev/doc/diagnostics    
- Use containers with limited resources to test
- The challenge was to use a CSV file as output, but in the future could use a database as output

---
## Extra
### Converting CSV result into HTML
The app has the capability to convert the CSV result into a static HTML file, to give a visual prefiew of the imagens and colors.
By default its already enabled. You can change the configs in the `.Env` file modifying the following:
```shell
GENERATE_HTML=true
HTML_TEMPLATE_FILENAME=./web/templates/result.tmpl
HTML_OUTPUT_FILENAME=index.html
```
When enabled, the HTML page should open automatically in the browser after the program ends.