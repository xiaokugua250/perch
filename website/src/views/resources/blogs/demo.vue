<template>
  <div>
    <head>
      <meta charset="utf-8">
      <meta http-equiv="X-UA-Compatible" content="IE=edge">
      <meta name="viewport" content="width=device-width, initial-scale=1">
      <title>Blog - Free Bulma template</title>
      <link rel="icon" type="image/png" sizes="32x32" href="../images/favicon.png">
      <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
      <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/overlayscrollbars/1.9.1/css/OverlayScrollbars.min.css">
      <link href="https://fonts.googleapis.com/css?family=Open+Sans" rel="stylesheet">
      <!-- Bulma Version 0.9.0-->
      <link rel="stylesheet" href="https://unpkg.com/bulma@0.9.0/css/bulma.min.css">
      <link rel="stylesheet" type="text/css" href="../../../styles/articles/css/blogs.css">
    </head>

    <body>
      <!-- START NAV -->
      <nav class="navbar">
        <div class="container">
          <div class="navbar-brand">
            <a class="navbar-item" href="../">
              <img src="../../../assets/images/resources/blogs/bulma.png" alt="Logo">
            </a>
            <span class="navbar-burger burger" data-target="navbarMenu">
              <span />
              <span />
              <span />
            </span>
          </div>
          <div id="navbarMenu" class="navbar-menu">
            <div class="navbar-end">
              <a class="navbar-item is-active">
                Home
              </a>
              <a class="navbar-item">
                Examples
              </a>
              <a class="navbar-item">
                Features
              </a>
              <a class="navbar-item">
                Team
              </a>
              <a class="navbar-item">
                Archives
              </a>
              <a class="navbar-item">
                Help
              </a>
              <div class="navbar-item has-dropdown is-hoverable">
                <a class="navbar-link">
                  Account
                </a>
                <div class="navbar-dropdown">
                  <a class="navbar-item">
                    Dashboard
                  </a>
                  <a class="navbar-item">
                    Profile
                  </a>
                  <a class="navbar-item">
                    Settings
                  </a>
                  <hr class="navbar-divider">
                  <div class="navbar-item">
                    Logout
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </nav>
      <!-- END NAV -->
      <section class="hero is-primary">
        <div class="hero-body">
          <div class="container">
            <h1 class="title">
              Primary title
            </h1>
            <h2 class="subtitle">
              Primary subtitle
            </h2>
          </div>
        </div>
      </section>
      <nav class="level">
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Tweets</p>
            <p class="title">3,456</p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Following</p>
            <p class="title">123</p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Followers</p>
            <p class="title">456K</p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Likes</p>
            <p class="title">789</p>
          </div>
        </div>
      </nav>
      <div class="container is-max-desktop">

        <div class="notification is-primary">
          This container has a <code>max-width</code> of <code>$desktop - $container-offset</code> on widescreen and fullhd.
        </div>
      </div>
      <section class="section">
        <div class="container">
          <h1 class="title">Section</h1>
          <h2 class="subtitle">
            A simple container to divide your page into <strong>sections</strong>, like the one you're currently reading
          </h2>
          <div class="bd-example">
            <div class="content is-large">
              <h1>Hello World</h1>
              <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla accumsan, metus ultrices eleifend gravida, nulla nunc varius lectus, nec rutrum justo nibh eu lectus. Ut vulputate semper dui. Fusce erat odio, sollicitudin vel erat vel, interdum mattis neque.</p>
              <h2>Second level</h2>
              <p>Curabitur accumsan turpis pharetra <strong>augue tincidunt</strong> blandit. Quisque condimentum maximus mi, sit amet commodo arcu rutrum id. Proin pretium urna vel cursus venenatis. Suspendisse potenti. Etiam mattis sem rhoncus lacus dapibus facilisis. Donec at dignissim dui. Ut et neque nisl.</p>
              <ul>
                <li>In fermentum leo eu lectus mollis, quis dictum mi aliquet.</li>
                <li>Morbi eu nulla lobortis, lobortis est in, fringilla felis.</li>
                <li>Aliquam nec felis in sapien venenatis viverra fermentum nec lectus.</li>
                <li>Ut non enim metus.</li>
              </ul>
            </div>
          </div>
        </div>

      </section>

      <article class="media ">
        <!---
  <figure class="media-left">
    <p class="image is-64x64">
      <img src="https://bulma.io/images/placeholders/128x128.png">
    </p>
  </figure>
  --->
        <div class="media-content">
          <div class="field">
            <p class="control">
              <textarea class="textarea" placeholder="Add a comment..." />
            </p>
          </div>
          <nav class="level">
            <div class="level-left">
              <div class="level-item">
                <a class="button is-info">Submit</a>
              </div>
            </div>
            <div class="level-right">
              <div class="level-item">
                <label class="checkbox">
                  <input type="checkbox"> Press enter to submit
                </label>
              </div>
            </div>
          </nav>
        </div>
      </article>
      <footer class="footer">
        <div class="content has-text-centered">
          <p>
            <strong>Bulma</strong> by <a href="https://jgthms.com">Jeremy Thomas</a>. The source code is licensed
            <a href="http://opensource.org/licenses/mit-license.php">MIT</a>. The website content
            is licensed <a href="http://creativecommons.org/licenses/by-nc-sa/4.0/">CC BY NC SA 4.0</a>.
          </p>
        </div>
      </footer>
    </body>
  </div>
</template>
<!-- <script src="../js/bulma.js"></script> -->
<script async type="text/javascript" src="../js/bulma.js"></script>
<script src='https://cdnjs.cloudflare.com/ajax/libs/overlayscrollbars/1.9.1/js/OverlayScrollbars.min.js'></script>
<script>
document.addEventListener("DOMContentLoaded", function() {
  //The first argument are the elements to which the plugin shall be initialized
  //The second argument has to be at least a empty object or a object with your desired options
  OverlayScrollbars(document.querySelectorAll("body"), { });
});
</script>
