<template>
  <div>

    <head>
      <meta charset="UTF-8">
      <meta name="viewport" content="width=device-width, initial-scale=1.0">
      <meta http-equiv="X-UA-Compatible" content="ie=edge">
      <title>站长简介:苦瓜先生</title>
      <!-- <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css"> -->
      <link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.4.1/css/all.css" integrity="sha384-5sAR7xN1Nv6T6+dT2mhtzEpVJvfS3NScPQTrOxhwjIuvcA67KV2R5Jz6kr4abQsz" crossorigin="anonymous">
      <link href="https://fonts.googleapis.com/css?family=Open+Sans:300,400,700" rel="stylesheet">

    </head>

    <body>
      <!-- Navigation bar -->
      <nav class="navbar is-primary is-fixed-top">
        <div class="navbar-brand">
          <div class="navbar-burger burger" data-target="navbarExampleTransparentExample">
            <span />
            <span />
            <span />
          </div>
        </div>

        <div id="navbarExampleTransparentExample" class="navbar-menu">
          <div class="navbar-end">
            <a class="navbar-item" href="#about">
              <span class="icon">
                <i class="fas fa-info" />
              </span>
              <span>关于我</span>
            </a>
            <a class="navbar-item" href="#services">
              <span class="icon">
                <i class="fas fa-bars" />
              </span>
              <span>技术服务与技术咨询</span>
            </a>
            <a class="navbar-item" href="#resume">
              <span class="icon">
                <i class="fas fa-file-alt" />
              </span>
              <span>简历</span>
            </a>
            <a class="navbar-item" href="#portfolio">
              <span class="icon">
                <i class="fas fa-th-list" />
              </span>
              <span>个人项目</span>
            </a>
            <a class="navbar-item" href="#contact">
              <span class="icon">
                <i class="fas fa-envelope" />
              </span>
              <span>联系方式</span>
            </a>
          </div>
        </div>
      </nav>

      <!-- Header -->
      <section class="hero is-primary  is-fullheight-with-navbar">
        <div class="hero-body">
          <div class="container">Hello! I am
            <h1 class="title is-1">
              苦瓜先生
            </h1>
            <h2 class="subtitle is-3">
              全栈开发
              专注于Golang、CloudNative和DevOps...
            </h2>
            <h6 class=" is-3">
              提供关于计算机软件开发等方面的技术支持和技术咨询
            </h6>
          </div>
        </div>
      </section>

      <!-- About -->
      <section id="about" class="section">
        <!-- Title -->
        <div class="section-heading">
          <h3 class="title is-2">About Me</h3>
          <h4 class="subtitle is-5">苦瓜先生</h4>
          <div class="container">
            <p>后端开发人员，程序开发 <strong>4 年经验</strong> 对计算机软件开发、计算机网络和系统平台维护有较多经验，工作语言为
              <strong>Golang</strong>,对C、Python、Java 等编程语言均有过实际项目开发经验。目前工作重心为<strong>云原生的PASS平台</strong>建设和构建，并且最近项目为<strong>基于Kubernetes的容器平台</strong>开发和维护。与此同时，对Devops、CI/CD等方面也较为熟悉，并有过实际落地项目。
              .</p>
          </div>
        </div>

        <div class="columns has-same-height is-gapless">
          <div class="column">
            <!-- Profile -->
            <div class="card">
              <div class="card-content">
                <h3 class="title is-4">Profile</h3>

                <div class="content">
                  <table class="table-profile">
                    <tr>
                      <th colspan="1" />
                      <th colspan="2" />
                    </tr>
                    <tr>
                      <td>Address:</td>
                      <td>广东-广州</td>
                    </tr>
                    <!--                    <tr>-->
                    <!--                      <td>Phone:</td>-->
                    <!--                      <td>0123-456789</td>-->
                    <!--                    </tr>-->
                    <tr>
                      <td>Email:</td>
                      <td>liangdu1992@gmail.com</td>
                    </tr>
                  </table>
                </div>
                <br>
                <div class="buttons has-addons is-centered">
                  <a href="https://github.com/xiaokugua250" class="button is-link">Github</a>
                  <!--  <a href="#" class="button is-link">LinkedIn</a>-->
                  <a href="https://twitter.com/xiaokugua250" class="button is-link">Twitter</a>
                  <!--       <a href="#" class="button is-link">CodeTrace</a>-->
                </div>
              </div>
            </div>
          </div>
          <div class="column">
            <!-- Profile picture -->
            <div class="card">
              <div class="card-image">
                <figure class="image is-4by4">
                  <img src="../../../assets/images/sites/users/profiles/twitter.png" alt="Placeholder image">
                </figure>
              </div>
            </div>
          </div>
          <div class="column">
            <!-- Skills -->
            <div class="card">
              <div class="card-content skills-content">
                <h3 class="title is-4">Skills</h3>
                <div class="content">

                  <article class="media">
                    <div class="media-content">
                      <div class="content">
                        <p>
                          <strong>Golang:</strong>
                          <br>
                          <progress class="progress is-primary" value="90" max="100" />
                        </p>
                      </div>
                    </div>
                  </article>

                  <article class="media">
                    <div class="media-content">
                      <div class="content">
                        <p>
                          <strong>Kubernetes/Container:</strong>
                          <br>
                          <progress class="progress is-primary" value="90" max="100" />
                        </p>
                      </div>
                    </div>
                  </article>

                  <article class="media">
                    <div class="media-content">
                      <div class="content">
                        <p>
                          <strong>Linux&Databases&C&Python:</strong>
                          <br>
                          <progress class="progress is-primary" value="75" max="100" />
                        </p>
                      </div>
                    </div>
                  </article>

                  <article class="media">
                    <div class="media-content">
                      <div class="content">
                        <p>
                          <strong>DevOps</strong>
                          <br>
                          <progress class="progress is-primary" value="95" max="100" />
                        </p>
                      </div>
                    </div>
                  </article>

                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      <!-- Skills tags -->
      <!-- <br>
          <div class="tags custom-tags">
            <span class="tag is-light">Node.js</span><span class="tag is-light">Express.js</span><span class="tag is-light">VueJS</span><span
              class="tag is-light">JavaScript</span><span class="tag is-light">HTML5</span><span class="tag is-light">Canvas</span><span
              class="tag is-light">CSS3</span><span class="tag is-light">Bulma</span><span class="tag is-light">Bootstrap</span><span
              class="tag is-light">jQuery</span><span class="tag is-light">Pug</span><span class="tag is-light">Stylus</span><span
              class="tag is-light">SASS/SCSS</span><span class="tag is-light">Webpack</span><span class="tag is-light">Git</span><span
              class="tag is-light">ASP.NET Web Forms</span><span class="tag is-light">MSSQL</span><span class="tag is-light">MongoDB</span><span
              class="tag is-light">Apache Cordova</span><span class="tag is-light">Chrome Extensions</span>
          </div> -->

      <!-- Services -->
      <section id="services" class="section">
        <div class="section-heading">
          <h3 class="title is-2">Services</h3>
          <h4 class="subtitle is-5">技术服务与技术咨询</h4>
        </div>
        <div class="container">
          <div class="columns">
            <div class="column">
              <div class="box">
                <div class="content">
                  <h4 class="title is-5">软件系统开发</h4>
                  Develop Front End using latest standards with HTML5/CSS3 with added funtionality using JavaScript and
                  Vue.js.
                </div>
              </div>
            </div>
            <div class="column">
              <div class="box">
                <div class="content">
                  <h4 class="title is-5">技术运维服务</h4>
                  Develop Back End application/service using Node.js or ASP .NET and SQL server or Mongo DB databases.
                </div>
              </div>
            </div>
          </div>

          <div class="columns">
            <div class="column">
              <div class="box">
                <div class="content">
                  <h4 class="title is-5">技术交流与培训</h4>
                  Develop Front End using latest standards with HTML5/CSS3 with added funtionality using JavaScript and
                  Vue.js.
                </div>
              </div>
            </div>
            <div class="column">
              <div class="box">
                <div class="content">
                  <h4 class="title is-5">其他软件服务</h4>
                  Develop Back End application/service using Node.js or ASP .NET and SQL server or Mongo DB databases.
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      <!-- Resume -->
      <section id="resume" class="section">
        <div class="section-heading">
          <h3 class="title is-2">Resume</h3>
          <h4 class="subtitle is-5">工作经历</h4>
          <a href="#" class="button is-link is-medium">
            <span class="icon">
              <i class="fas fa-file-alt" />
            </span>
            <span>点击下载简历</span>
          </a>
        </div>
      </section>

      <!-- Portfolio -->
      <section id="portfolio" class="section">
        <div class="container">
          <div class="section-heading">
            <h3 class="title is-2">个人项目</h3>
            <h4 class="subtitle is-5">Github项目</h4>
          </div>
          <br>

          <div class="container portfolio-container">
            <div class="columns">
              <div class="column ">
                <!-- admin -->
                <div class="card">
                  <header class="card-header">
                    <p class="card-header-title">
                      <span>Perch <small>(Light)</small></span>
                      <span class="is-pulled-right">
                        <a href="https://github.com/jgthms/bulma/releases/tag/0.7.2">
                          <span class="tag is-default">v0.7.2</span>
                        </a>
                        <span class="tag is-default">Desktop</span>
                        <span class="tag is-danger">WIP</span>
                      </span>
                    </p>
                  </header>
                  <div class="card-content">
                    <figure class="image">
                      <img src="../../../assets/images/sites/users/profiles/perch_logo.jpg" alt="Admin template screenshot">
                    </figure>
                  </div>
                  <footer class="card-footer">
                    <a href="https://github.com/xiaokugua250/perch" class="card-footer-item">Preview</a>
                    <a href="https://github.com/xiaokugua250/perch" class="card-footer-item">Source
                      Code</a>
                  </footer>
                </div>
                <!-- cards -->

                <!-- forum -->

              </div>
              <div class="column is-4">

                <!-- hero -->
                <div class="card">
                  <header class="card-header">
                    <p class="card-header-title">
                      <span>Harbor-go</span>
                      <span class="is-pulled-right">
                        <a href="https://github.com/jgthms/bulma/releases/tag/0.7.2">
                          <span class="tag is-default">v0.7.2</span>
                        </a>

                      </span>
                    </p>
                  </header>
                  <div class="card-content">
                    <figure class="image">
                      <img src="../../../assets/images/sites/users/profiles/habor.png" alt="Hero template screenshot">
                    </figure>
                  </div>
                  <footer class="card-footer">
                    <a href="https://github.com/xiaokugua250/harbor-go" class="card-footer-item">Preview</a>
                    <a href="https://github.com/xiaokugua250/harbor-go" class="card-footer-item">Source
                      Code</a>
                  </footer>
                </div>
                <!--kanban -->
                <div class="card">
                  <header class="card-header">
                    <p class="card-header-title">
                      <span>Siniperca</span>
                      <span class="is-pulled-right" />
                    </p>
                  </header>
                  <div class="card-content">
                    <figure class="image">
                      <img src="../../../assets/images/sites/users/profiles/crd.jpeg" alt="harbor screenshot">
                    </figure>
                  </div>
                  <footer class="card-footer">
                    <a href="https://github.com/xiaokugua250/Siniperca" class="card-footer-item">Preview</a>
                    <a href="https://github.com/xiaokugua250/Siniperca" class="card-footer-item">Source
                      Code</a>
                  </footer>
                </div>
                <!-- login -->

              </div>

            </div>
          </div>
        </div>
      </section>

      <!-- Contact -->
      <section id="contact" class="section">
        <div class="container">
          <div class="section-heading">
            <h3 class="title is-2">Contact</h3>
            <h4 class="subtitle is-5">联系我</h4>
          </div>
          <br>

          <div class="columns">
            <div class="column is-6 is-offset-3">
              <div class="box">
                <div class="field">
                  <label class="label">Name</label>
                  <div class="control">
                    <input class="input" type="text" placeholder="Text input">
                  </div>
                </div>

                <div class="field">
                  <label class="label">Email</label>
                  <div class="control has-icons-left">
                    <input class="input" type="email" placeholder="Email input" value="">
                    <span class="icon is-small is-left">
                      <i class="fas fa-envelope" />
                    </span>
                  </div>
                </div>

                <div class="field">
                  <label class="label">Message</label>
                  <div class="control">
                    <textarea class="textarea" placeholder="Textarea" />
                  </div>
                </div>

                <div class="field is-grouped has-text-centered">
                  <div class="control">
                    <button class="button is-link is-large"><span class="icon">
                                                              <i class="fas fa-envelope" />
                                                            </span>
                      <span>Submit</span></button>
                  </div>
                </div>
              </div>
            </div>
          </div>

        </div>
      </section>

      <!-- Footer -->
      <footer class="footer">
        <div class="section-heading">
          <p>
            <strong>Bulma Personal Template</strong> by <a href="https://mubaidr.js.org">Muhammad Ubaid Raza</a>. The
            source code is licensed
            <a href="http://opensource.org/licenses/mit-license.php">MIT</a>.
          </p>
        </div>
      </footer>
    </body>
  </div>
</template>
