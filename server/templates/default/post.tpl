{{ block "header.tpl" . }}{{ end }}
<link rel="stylesheet" type="text/css" href="/themes/assets/markdown/css/github-markdown.css">
<link rel="stylesheet" type="text/css" href="/themes/assets/css/highlight.css">

<script src="/themes/assets/js/marked.min.js"></script>
<script src="/themes/assets/js/highlight.js"></script>
{{ block "nav.tpl" . }}{{ end }}
<div class="container">
  <main class="main">
    <article class="post">
        <h1 class="title">{{ .Data.Post.PrimaryPost.Title }}</h1>
        <div class="post-footer">
          <div class="left">
            <span>{{ .Data.Post.PrimaryPost.Author}}</span>
            <span>发布@{{func_format_timestamp .Data.Post.PrimaryPost.CreatedAt "date"}}</span>
            <span>浏览量({{ .Data.Post.PrimaryPost.VisitCount}})</span>
            <span>点赞({{ .Data.Post.PrimaryPost.LikeCount}})</span>
          </div>
        </div>
        <div class="content markdown-body" id="content-container">{{ .Data.Post.ExtraPost.Content }}</div>
        <div class="pre-next">
            <div class="pre">
                上一篇：<a href="/post?id=1111">aaaaa</a>
              </div>
              <div class="next">
                下一篇：<a href="/post?id=22222">bbbbb</a>
              </div>
        </div>
      </article>
  </main>
</div>
<script>
    const contentEle = document.getElementById('content-container');
    hljs.highlightAll();

    marked.setOptions({
        highlight: function (code, language) {
            const validLanguage = hljs.getLanguage(language) ? language : 'javascript';
            return hljs.highlight(code, {language: validLanguage}).value;
        },
    });
    contentEle.innerHTML = marked.parse(contentEle.innerHTML);
</script>
{{ block "footer.tpl" . }}{{ end }}