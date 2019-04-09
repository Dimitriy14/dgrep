<!DOCTYPE html>
<html>
<head>
    <title>Statistic</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link href="/static/css/main.css" rel="stylesheet">
</head>
<body>
<h2 class = "stathead">{{call .i18n "Get_statistics"}}</h2>
    <div class="block-container">
        <ul>
            {{range $_,$field := .Statistic}}
                <li>{{$field}}:</li>
            {{end}}
        </ul>
        {{range $_, $val := .Stat}}
            <ul>
                {{range $_,$stat := $val}}
                    <li>{{$stat}}</li>
                {{end}}
            </ul>     
        {{end}}
    </div> 
    <div class="button-center"><a href="/" class="back">{{call .i18n "Back"}}</a></div>
</body>
</html>