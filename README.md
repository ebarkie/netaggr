# Network aggregator/summarizer

This tool takes a list of networks and attempts to reduce it by removing smaller networks
which are already represented by larger networks and joining adjacent networks to form
larger networks.  The input and output represent exactly the same set of IP addresses.

Networks may be formatted as IPv4/6 CIDR or an IPv4 address and a dot-decimal subnet mask.

## Algorithm

The algorithm works as follows:

1. Sort the network list in ascending order
   - Join the network and mask bytes and compare lexicographically
2. Assimilate smaller networks into larger networks
   - Iterate the network list from the first network to the second-to-last network
      - If the next network falls within the current network, then delete the next network
3. Aggregate adjacent networks to form larger networks
   - Iterate the network list from the first network to the second-to-last network
      - If the next network decremented by one falls within the current network, and the
        masks/prefix lengths are equal
         - Decrement the prefix length of the current network by one
         - Delete the next network
         - Decrement the iteration index by one

## Usage

```
Usage of ./netaggr:
  -aggr
    	perform network aggregation (default true)
  -assim
    	perform network assimilation (default true)
  -in string
    	input file
  -notation string
    	output notation: "cidr" or "dd" (default "cidr")
```

If the `-in` flag is not specified it will read from stdin. To turn off assimilation or
aggregation use the `-assim=false` and `-aggr=false` flags.

## License

Copyright (c) 2018-2019 Eric Barkie. All rights reserved.  
Use of this source code is governed by the MIT license
that can be found in the [LICENSE](LICENSE) file.
