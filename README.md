# calc_go
> ✨
> It`s a web service, based on math expressions using numbers, signs as "+", "-", "*" and "/".
> which calculates the value of the resulting expression✨<br/>
> <h3>Calculator supports next operations:<h3>
> <h4> Add "+"<br/>
> Subtraction "-" <br/>
>Multiply * <br/>
> Division "/" <br/>
> Brackets "()"

>[!IMPORTANT]HTTP Status codes, that provides calculator:
<table>
  </thead>
  <tbody>
    <tr>
      <td>OK</td>
      <td>200</td>
      <td>POST</td>
      <td>{ "result":"expression result"}</td>
    </tr>
    <tr>
      <td>Calculation error</td>
      <td>400</td>
      <td>POST</td>
      <td>{"error":"calculation error, incorrect expression"}</td>
    </tr>
    <tr>
      <td>Wrong Method</td>
      <td>405</td>
      <td>any (exc POST)</td>
      <td>{"error":"Method is not allowed"}</td>
    </tr>
    <tr>
      <td>Invalid expression</td>
      <td>422</td>
      <td>POST</td>
      <td>{"error":"Expression is not valid"}</td>
    </tr>
    <tr>
      <td>Internal server error</td>
      <td>500</td>
      <td>POST</td>
      <td>{error: "Internal server error"}</td>
</tr>
</tbody></table><br>

>[!IMPORTANT] To start program you need to write this these lines to the terminal:
 <h3> go run cmd/main/main.go

<h2> Examples: </h2>

>[!NOTE] 
> 200 OK<br>
> <h4>curl 'localhost:8000/api/v1/calculate' \
>--header 'Content-Type: application/json' \
>--data '{"expression":"1+1"}' <h4>
> 
Result:
> <h4>{"result":2} <h4>

>400 (Bad Request) <h4>curl 'localhost:8000/api/v1/calculate' \
>--header 'Content-Type: application/json' \
>--data 'feijhgoerhn' <h4>
>
Result: 
>{"error":"calculation error, incorrect expression"}

> 405 (Wrong Method) 
> <h4>curl --request GET 'localhost:4200/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression":"1+1"}'<h4>

Result:
> {"error":"Method is not allowed"}

> 422 (Unprocessable Entity)
<h4>curl 'localhost:4200/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression":"1+"}' <h4>

Result:
>{"error":"Expression is not valid"}