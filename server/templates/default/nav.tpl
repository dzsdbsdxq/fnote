<nav class="nav">
<ul>
  <li><a href="/" class="nav-item">首页</a></li>
  {{ range $key,$value := func_get_menus }}
    {{if $value.ShowInNav}}
        <li><a href="/categories/{{$value.Route}}" class="nav-item {{if eq $.PageIdx $value.Route}}active{{else}}{{end}}">{{$value.Name}}</a></li>
    {{end}}
  {{end}}
</ul>
<div class="search">
  <input type="search" class="ui-input" placeholder="搜索" autofocus />
</div>
</nav>