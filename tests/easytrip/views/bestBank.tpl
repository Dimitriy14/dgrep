<div class = "mycontainer">

<h1>{{call .i18n "Best_Banks"}}</h1>
<h2>{{call .i18n .TitleBuy }}</h2>
{{if $self := .}}
{{range .Buy}}
<div class="bank">
<ul>
<h2 class="bank-name">{{call $self.i18n .BankName}}</h2>
<li>
{{.CodeAlpha}}: {{.RateBuy}}
</li>
<br>
</ul>
</div>
{{end}}
{{end}}

<h2>{{call .i18n .TitleSale }}</h1>
{{if $self := .}}
{{range .Sale}}
<div class="bank">
<ul>
<h2 class="bank-name">{{call $self.i18n .BankName}}</h2>
<li>
{{.CodeAlpha}}: {{.RateSale}}
</li>
<br>
</ul>
</div>
{{end}}
{{end}}
<a href="/" class="back">{{call .i18n "Back"}}</a>
</div>