<div>History</div>

{{range $key, $val := .Req}}
   <div><a href = "{{$val.Link}}">{{$val.Start}}</a></div>
{{end}}