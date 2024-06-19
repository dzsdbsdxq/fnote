<!DOCTYPE html>
<html lang="zh-CN">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta name="referrer" content="no-referrer" />
    <title>{{.Data.BlogName}}</title>
    <meta name="keywords" content="">
    <link rel="stylesheet" href="/themes/assets/css/lulu/ui.min.css" />
    <link rel="stylesheet" href="/themes/assets/css/style.css">
  </head>
  <body>
    <header class="header">
      <div class="nickname">{{.Data.Meta.WebSiteConfig.WebsiteName}}</div>
      <div class="sign">{{.Data.Meta.WebSiteConfig.SocialInfoConfig}}</div>
    </header>
    <div class="container">
      <nav class="nav">
        <ul>
          <li><a href="/" class="nav-item">首页</a></li>
          {{ range $key,$value:=$.Data.Menus }}
          <li><a href="/post?id=1" class="nav-item">{{$value.Name}}</a></li>
          {{end}}
        </ul>
        <div class="search">
          <input type="search" class="ui-input" placeholder="搜索" autofocus />
        </div>
      </nav>
      <main class="main">
        <div class="list">
            <ul>
                <li>
                    <a href="/post?id=1111">
                        <span class="red">[置顶]</span>
                        <span>走出人格陷阱</span>
                    </a>
                    <hr class="hr" />
                    <time class="date">2024/05/17</time>
                </li>
                <li>
                    <a href="/post?id=1111">
                        <span style="display:none" class="red">[置顶]</span>
                        <span>Kind快速部署Kubernetes多集群</span>
                    </a>
                    <hr class="hr" />
                    <time class="date">2024/05/17</time>
                </li>
                <li>
                    <a href="/post?id=1111">
                        <span style="display:none" class="red">[置顶]</span>
                        <span>Ubuntu 使用Docker compose快速部署Zabbix 6.4</span>
                    </a>
                    <hr class="hr" />
                    <time class="date">2024/05/17</time>
                </li>
            </ul>
            <div id="page" class="page" data-total="100"></div>
          </div>
      </main>
    </div>
    <footer class="footer">
      <span class="icp">
      粤ICP-11111 由 <a href="#" target="_blank">fnote</a>  强力驱动 | 主题 - Yun v0.0.1 Powered by
        <a href="#" target="_blank">李守杰</a>
        </span>
    </footer>
  </body>
</html>