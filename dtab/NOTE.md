# Dtab
## Names

## Paths
Names beginning with character '/' are hierarchical paths. Abstraction location.
Path names should be bound by current namespace

## Delegation tables
Define namespace
Syntax

```
src => dest
```

To rewrite all prefix

```
/s => /s#/foo/bar
/s/crawler => /s#/foo/bar/crawler
```

Note: match on "path" component.

Could match wildcard (*) 

```
/s#/*/bar => /t/bah
make
/s#/foo/bar/baz or /s#/boo/bar/baz
to /t/bah/baz
```

Path begin with '/$/' called "system path". It needs predefine Namers.
For example: /$/namer/path.. use given "Namers" to resolve

Dtab may contain "#" (comment?) and split into multiple lines.

```
# delegation for /s
/s => /a      # prefer /a
    | ( /b    # or share traffic between /b and /c
      & /c
      );
equivalent to
/s => /a | (/b & /c);
```

First rule choosed, using fallback mechanism

You can specify any number of alternates 
```
/humphrys | /smitten | /birite | /three-twins …
```

Dtabs also support unions with the following syntax /iceCreamStore => /humphrys & /smitten. 
In this example we have an equal chance of routing the path to either store. 
If we wanted to be more likely to enter one store than another, we can add weights to each path:

```
/smitten       => 3 * /SF/Octavia/432 & 1 * /SF/California/2404;
/iceCreamStore => 0.7 * /humphrys & 0.3 * /smitten;
```

## Namer
Context bound look up method?