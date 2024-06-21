{{ block "header.tpl" . }}{{ end }}
{{ block "nav.tpl" . }}{{ end }}
<div class="container">
  <main class="main">
    <div class="list">
        <ul>
            {{ range $key,$value := func_posts .PageRequest.PageNo .PageRequest.PageSize "created_at" "DESC" .PageIdx "css"}}
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