{{ block "header.tpl" . }}{{ end }}
    <div class="container">
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
      <main class="main">
        <div class="list">
            <ul>
                {{ range $key,$value := func_latest_posts 10}}
                <li>
                    <a href="/post/{{$value.PrimaryPost.Id}}">
                        {{if eq $value.PrimaryPost.StickyWeight 1}}
                        <span class="red">[置顶]</span>
                        {{end}}
                        <span>{{$value.PrimaryPost.Title}}</span>
                    </a>
                    <hr class="hr" />
                    <time class="date">{{func_format_timestamp $value.PrimaryPost.CreatedAt "date"}}</time>
                </li>
                {{end}}
            </ul>
            <div id="page" class="page" data-total="100"></div>
          </div>
      </main>
    </div>
{{ block "footer.tpl" . }}{{ end }}