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

      <section class="hero is-info is-medium is-bold">
        <div class="hero-body">
          <div class="container has-text-centered">
            <h1 class="title">Lorem ipsum dolor sit amet, consectetur adipiscing elit, <br>sed eiusmod tempor incididunt ut labore et dolore magna aliqua</h1>
          </div>
        </div>
      </section>

      <div class="container">
        <!-- START ARTICLE FEED -->
        <section class="articles">

          <div class="tile is-ancestor column ">
            <div class="tile is-vertical is-7 is-offset-1 is-half">

              <div class="tile is-parent">
                <article class="tile is-child notification is-half">
                  <p class="title">Wide tile</p>
                  <p class="subtitle">Aligned with the right tile</p>
                  <div class="content">
                    <div class="control">

                    <textarea class="textarea is-primary" placeholder="Primary textarea" :value="input" @input="update" rows="50"></textarea>
                    </div>


                  </div>
                </article>
              </div>
            </div>
            <div class="tile is-parent">
              <article class="tile is-child notification is-white ">
                <div class="content">
                  <p class="title">Tall tile</p>
                  <p class="subtitle">With even more content</p>
                  <div class="content">
                    <div v-html="compiledMarkdown"></div>
                  </div>
                </div>
              </article>
            </div>
          </div>
          <div class="field is-grouped">
            <div class="control">
              <button class="button is-link" @click="userCreateBlog">Submit</button>
            </div>
            <div class="control">
              <button class="button is-link is-light">Cancel</button>
            </div>
          </div>
        </section>
      <!-- END ARTICLE FEED -->
      </div>

    </body>
  </div>
</template>
<!-- <script src="../js/bulma.js"></script> -->

<script>
import {createblogs, getBlogs} from "@/api/resources-blogs";
import marked  from 'marked'
import _ from 'lodash'
export default {
  data() {
    return {
      // data,
      id: 0,
      blog:null,
      blogs: [],
      input: "# hello",
      showDetailIcon:false,
      listQuery: {
        page: 1,
        limit: 20,
        importance: undefined,
        title: undefined,
        type: undefined,
        sort: '+id'
      },
      defaultOpenedDetails: [1]

    }
  },
  computed: {
    compiledMarkdown: function() {
      return marked(this.input, { sanitize: true });
    }
  },
  created() {
    // this.getList()
    //this.resourceBlogsGet()
  },
  methods: {
   update: _.debounce(function(e) {
      this.input = e.target.value;
    }, 300),
    userCreateBlog:function (){
    console.log("===>",this.blog)
      this.listLoading = true
      createblogs(this.blog).then(response => {
        //        this.list = response.data.items
        this.blogs = response.spec
        // console.log('===>', this.blogs)
        this.total = response.total

        // Just to simulate the time of the request
        setTimeout(() => {
          this.listLoading = false
        }, 1.5 * 1000)
      })
    }
    }


}
</script>
