# urrls/mold
#The ullr/mold package implements support for dynamic JSON keys and pattern-matched object structure transformation

## Install
    go get github.com/urrls/mold

***

#### Dynamic key entry syntax
{ "user_id=-id": 25 }
#If the name of the key in the source matches the name of the key in the sample, a simplified notation is available
{ "id=-id": 25 }
equivalent declaration
{ "id=-": 25 }

#### Examples
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
***

#### Valid Key Operators
##### "=-" Passes a value to the form only if it is found in the source, matches the value type, and is on the same level as the form, otherwise leaves the form's value. Child objects are filled recursively and only dynamic values.
##### "=!" #Passes a value to the form only if it is found in the source, matches the value type, and is at the same level as the form, otherwise null is passed. Child objects are filled recursively and only dynamic values.
##### "==" Passes a value to the form, regardless of types. Child objects are passed in their entirety, without recursive parsing. Recommended if you are expecting a primitive of unknown type in the value.

##### "<-" Search for a value by key throughout the source, including child and parent objects. Passes a value to the form, regardless of types. Child objects and arrays are also processed to find dynamic keys. If there is no value in the source, a value is returned in the form
##### "<!" Search for a value by key throughout the source, including child and parent objects. Passes a value to the form, regardless of types. Child objects and arrays are also processed to find dynamic keys. If there is no value in the source, null is returned.
##### "<<" Search for a value by key throughout the source, including child and parent objects. Passes a value to the form, regardless of types. Child objects and arrays are passed in full, without recursive parsing. It is recommended to use if you are expecting a primitive of an unknown type in the value.
***
