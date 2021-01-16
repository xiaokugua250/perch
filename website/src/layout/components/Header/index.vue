<template>
  <div >
    <section class="hero is-info is-medium is-bold">
      <div class="hero-head">
        <nav class="navbar">
          <div class="container">
            <div class="navbar-brand"  height="400">
              <a class="navbar-item" href="../" >
                <img src="../../../assets/images/sites/basic/z-gour.com.png" alt="Logo" >
              </a>
              <span class="navbar-burger burger" data-target="navbarMenu">
                  <span />
                  <span />
                  <span />
                </span>
            </div>
            <div id="navbarMenu" class="navbar-menu">
              <div class="navbar-end">
                <div class="tabs is-right">
                  <ul>
                    <li class="is-active"><a>Home</a></li>
                    <li><a href="">Tech&Lives</a></li>
                    <li><a href="">News</a></li>
                    <li><a href="">Team</a></li>

                    <li><a href="#/resources">Resources</a></li>
                    <li><a href="">Help</a></li>
                    <li><a href="#/search">Search</a></li>
                  </ul>
                  <span class="navbar-item">
                      <a class="button is-white is-outlined" href="https://github.com/BulmaTemplates/bulma-templates/blob/master/templates/hero.html">
                        <span class="icon">
                          <i class="fa fa-github" />
                        </span>
                        <span title="Hello from the other side">View Source</span>
                      </a>
                    </span>
                </div>
              </div>
            </div>
          </div>
        </nav>
      </div>
    </section>

  </div>
</template>


<script>
  export default {
    name: 'Header',
    /*computed: {
      cachedViews() {
        return this.$store.state.tagsView.cachedViews
      },
      key() {
        return this.$route.path
      }
    }*/
  }
</script>
