## ullrs/mold
The ullrs/mold package implements support for dynamic JSON keys and pattern-matched object structure transformation

### Install
    go get github.com/ullrs/mold

***

### Dynamic key entry syntax
{ "user_id=-id": 25 }
#If the name of the key in the source matches the name of the key in the sample, a simplified notation is available
{ "id=-id": 25 }
equivalent declaration
{ "id=-": 25 }

### Examples
```go
    sourceJson := `{ "name": "John", "experience": "5 years", "job": {"profession": "programmer", "job_title": "administrator"}, "gadgets": ["tablet", "phone", "laptop"] } `
    moldJson := `{ "first_name<-name": "not indicated", "job<-": {"stack<-profession": "Trainee", "job_title": "administrator", "experience<-": "Without experience"}, "gadgets<-gadgets": ["No found"] } `

    result, err := mold.Fill([]byte(sourceJson), []byte(moldJson))
    if err != nil {
	panic(err)
    }
    fmt.Println(string(result))
```
#### Result:
    >>>{"first_name":"John","gadgets":["tablet","phone","laptop"],"job":{"job_title":"administrator","stack":"programmer"}}%

> Pay attention to the rules for filling arrays: how the sample is used the first element of the mold array and all elements of the source array are compared with it
***

### Valid Key Operators
```go
<-, <!, <<, =-, =!, ==
```
##### "=-" Passes a value to the form only if it is found in the source, matches the value type, and is on the same level as the form, otherwise leaves the form's value. Child objects are filled recursively and only dynamic values.
```go
{ "name=-": "...", "last_name=-second_name": "..." }
```
##### "=!" #Passes a value to the form only if it is found in the source, matches the value type, and is at the same level as the form, otherwise null is passed. Child objects are filled recursively and only dynamic values.
```go
{ "name=!": "...", "last_name=!second_name": "..." }
```
##### "==" Passes a value to the form, regardless of types. Child objects are passed in their entirety, without recursive parsing. Recommended if you are expecting a primitive of unknown type in the value.
```go
{ "name==": "...", "last_name==second_name": "..." }
```
##### "<-" Search for a value by key throughout the source, including child and parent objects. Passes a value to the form, regardless of types. Child objects and arrays are also processed to find dynamic keys. If there is no value in the source, a value is returned in the form
```go
{ "name<-": "...", "last_name<-second_name": "..." }
```
##### "<!" Search for a value by key throughout the source, including child and parent objects. Passes a value to the form, regardless of types. Child objects and arrays are also processed to find dynamic keys. If there is no value in the source, null is returned.
```go
{ "name<!": "...", "last_name<!second_name": "..." }
```
##### "<<" Search for a value by key throughout the source, including child and parent objects. Passes a value to the form, regardless of types. Child objects and arrays are passed in full, without recursive parsing. It is recommended to use if you are expecting a primitive of an unknown type in the value.
```go
{ "name<<": "...", "last_name<<second_name": "..." }
```
***
