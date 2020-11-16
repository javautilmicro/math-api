### math-api
assignment 2: math api


The following endpoints can be executed via Postman
and these helper JSON stubs.



* http:/localhost:8088/min
* http:/localhost:8088/max
* http:/localhost:8088/avg
* http:/localhost:8088/median
* http:/localhost:8088/percentile


######Here is some example JSON, you can alter it whatever way you like, but to get a return value the only rule is that your list of numbers must have the same 'size' as the 'numberItems' property.

```
{
	"numberItems": 7,
	"numbers": [11, 35, 68, 5, 22, 34, 54]
}
```