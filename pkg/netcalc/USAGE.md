# netcalc
```
import "github.com/ebarkie/netaggr/pkg/netcalc"
```

Package netcalc performs calculations against IP networks.

It can parse networks formatted as IPv4/6 CIDR or an IPv4 address and a
dot-decimal subnet mask, and assimilate or aggregate them.

## Usage

#### func  Compare

```go
func Compare(a, b net.IPNet) int
```
Compare returns an integer comparing two IPNet's lexicographically. The result
will be 0 if a==b, -1 if a < b, and +1 if a > b.

#### func  DD

```go
func DD(n net.IPNet) string
```
DD returns the IP network n as a string formatted as an IPv4 address and a
dot-decimal subnet mask.

#### type Nets

```go
type Nets []*net.IPNet
```

Nets is a sorted slice of IPNet's. If this is populated by means other than
Parse then the caller is responsible for sorting.

#### func  Diff

```go
func Diff(a, b Nets) (added, deleted Nets)
```
Diff finds the differences between two slices of networks. It returns a slice of
added networks that don't exist in a but do in b, and a slice of deleted
networks that exist in a but don't in b.

#### func  Parse

```go
func Parse(r io.Reader) (Nets, error)
```
Parse parses single addresses or networks formatted as IPv4/6 addresses, IPv4/6
CIDR, or an IPv4 address and a dot-decimal subnet mask, like:

    192.0.2.1
    192.0.2.0/24
    192.0.2.0 255.255.255.0
    192.0.2.0/255.255.255.0

It returns a sorted list of Nets.

#### func (*Nets) Aggr

```go
func (nets *Nets) Aggr()
```
Aggr aggregates networks by joining adjacent networks to form larger networks.

#### func (*Nets) Assim

```go
func (nets *Nets) Assim()
```
Assim assimilates networks by removing smaller networks that are inside larger
networks.

#### func (Nets) Len

```go
func (n Nets) Len() int
```
sort.Interface implementation.

#### func (Nets) Less

```go
func (n Nets) Less(i, j int) bool
```

#### func (Nets) String

```go
func (n Nets) String() string
```

#### func (Nets) Swap

```go
func (n Nets) Swap(i, j int)
```
