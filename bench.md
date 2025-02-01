# Simple Benchmark

CPU: M3 Max(16 core)

## Encode

| tool        | user | sys | tot  | rate     | size | rate ratio | size ratio |  CPU% |
|:-----------:|:----:|:---:|:----:|:--------:|:----:|:----------:|:----------:|:-----:|
| gzip(fast)  | 16.4 | 0.4 | 16.8 | 387 MB/s | 984M | (1.0x)     | (100%)     |  100% |
| gzip(best)  | 225. | 0.7 | 225. |  29 MB/s | 766M |  0.1x      |   78%      |  100% |
| pgzip(best) | 367. | 1.7 | 24.3 | 269 MB/s | 740M |  0.7x      |   75%      | 1517% |
| pgzip       | 20.4 | 0.7 |  2.0 | 3.2 GB/s | 840M |  8.3x      |   85%      | 1055% |
| pgzip(fast) | 12.0 | 0.6 |  1.9 | 3.5 GB/s | 936M |  9.0x      |   95%      |  663% |
| pgzip(huff) |  8.2 | 0.8 |  1.7 | 3.9 GB/s | 3.4G | 10.1x      |  346%      |  529% |

## Decode

| tool                | user | sys | tot  | rate     | rate ratio |
|:-------------------:|:----:|:---:|:----:|:--------:|:----------:|
| gzip(fast) -> zcat  |  3.0 | 0.2 |  3.2 | 2.1 GB/s | (1.0x)     |
| gzip(fast) -> pzcat | 11.4 | 0.7 | 10.7 | 607 MB/s |  0.3x      |
| gzip(best) -> zcat  |  3.2 | 0.2 |  3.5 | 1.9 GB/s |  0.9x      |
| gzip(best) -> pzcat |  9.9 | 0.6 |  9.2 | 710 MB/s |  0.3x      |
| pgzip(best)         |  9.4 | 0.7 |  8.6 | 754 MB/s |  0.4x      |
| pgzip(best) -> zcat |  3.3 | 0.2 |  3.5 | 1.9 GB/s |  0.9x      |
| pgzip               | 10.3 | 0.7 |  9.6 | 676 MB/s |  0.3x      |
| pgzip       -> zcat |  3.5 | 0.2 |  3.7 | 1.8 GB/s |  0.9x      |
| pgzip(fast)         | 10.7 | 0.7 | 10.0 | 648 MB/s |  0.3x      |
| pgzip(fast) -> zcat |  3.5 | 0.2 |  3.7 | 1.7 GB/s |  0.8x      |
| pgzip(huff)         | 33.9 | 1.1 | 33.6 | 193 MB/s |  0.1x      |
| pgzip(huff) -> zcat | 12.0 | 0.3 | 12.3 | 530 MB/s |  0.3x      |
