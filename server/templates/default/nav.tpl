<nav class="nav">
<ul>
  <li><a href="/" class="nav-item">首页</a></li>
  {{ range $key,$value := func_get_menus }}
    {{if $value.ShowInNav}}
        <li><a href="/post?id=1" class="nav-item">{{$value.Name}}</a></li>
    {{end}}
  {{end}}
</ul>
<div class="search">
  <input type="search" class="ui-input" placeholder="搜索" autofocus />
</div>
</nav>