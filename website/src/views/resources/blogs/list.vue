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
      <!-- Bulma Version 0.9.0-->
      <link rel="stylesheet" href="https://unpkg.com/bulma@0.9.0/css/bulma.min.css">

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
      <section>
        <div class="container">
          <div class="notification is-primary">
            This container is <strong>centered</strong> on desktop and larger viewports.

          </div>
          <div v-for="blog in blogs">
            <div class="card ">
              <div class="card-image">
                <figure class="image">
                  <img src="../../../assets/images/resources/blogs/background_blogs.jpg" alt="Placeholder image" height="20" width="20">
                </figure>
                {{blog}}
              </div>
              <div class="card-content">
                <div class="media">
                  <div class="media-left">
                    <figure class="image is-48x48">
                      <img src="https://bulma.io/images/placeholders/96x96.png" alt="Placeholder image">
                    </figure>
                  </div>
                  <div class="media-content">
                    <p class="title is-4">John Smith</p>
                    <p class="subtitle is-6">@johnsmith</p>
                  </div>
                </div>

                <div class="content">
                  Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                  Phasellus nec iaculis mauris. <a>@bulmaio</a>.
                  <a href="#">#css</a> <a href="#">#responsive</a>
                  <br>
                  <time datetime="2016-1-1">11:09 PM - 1 Jan 2016</time>
                </div>
              </div>
            </div>
            <div>&nbsp;&nbsp;<br></br></div>
          </div>
        </div>
        <div class="container">
          <div class="notification is-primary">
            This container is <strong>centered</strong> on desktop and larger viewports.
          </div>
          <b-table
            ref="table"
            :data="blogs"
            paginated

            per-page="5"
            :opened-detailed="defaultOpenedDetails"
            detailed
            detail-key="id"
            aria-next-label="Next page"

            aria-previous-label="Previous page"
            aria-page-label="Page"
            aria-current-label="Current page"
            @details-open="(row) => $buefy.toast.open(`Expanded ${row.name}`)"
          >

            <b-table-column v-slot="props" field="id" label="ID" width="40" numeric>
              {{ props.row.id }}
            </b-table-column>

            <b-table-column v-slot="props" field="name" label="标题(Title)" sortable>
              <template v-if="showDetailIcon">
                {{ props.row.title }}
              </template>
              <template v-else>
                <a @click="props.toggleDetails(props.row)">
                  {{ props.row.title }}
                </a>
              </template>
            </b-table-column>

            <b-table-column v-slot="props" field="name" label="作者(Author)" sortable>
              <template v-if="showDetailIcon">
                {{ props.row.author }}
              </template>
              <template v-else>
                <a @click="props.toggleDetails(props.row)">
                  {{ props.row.author }}
                </a>
              </template>
            </b-table-column>
            <b-table-column v-slot="props" field="name" label="类别(Category)" sortable>
              <template v-if="showDetailIcon">
                {{ props.row.category }}
              </template>
              <template v-else>
                <a @click="props.toggleDetails(props.row)">
                  {{ props.row.category }}
                </a>
              </template>
            </b-table-column>
            <b-table-column v-slot="props" field="name" label="标签(tags)" sortable>
              <template v-if="showDetailIcon">
                {{ props.row.tags }}
              </template>
              <template v-else>
                <a @click="props.toggleDetails(props.row)">
                  {{ props.row.tags }}
                </a>
              </template>
            </b-table-column>
            <b-table-column v-slot="props" field="name" label="发布时间" sortable>
              <template v-if="showDetailIcon">
                {{ props.row.created_at }}
              </template>
              <template v-else>
                <a @click="props.toggleDetails(props.row)">
                  {{ props.row.created_at }}
                </a>
              </template>
            </b-table-column>
            <b-table-column v-slot="props" field="name" label="更新时间" sortable>
              <template v-if="showDetailIcon">
                {{ props.row.updated_at }}
              </template>
              <template v-else>
                <a @click="props.toggleDetails(props.row)">
                  {{ props.row.updated_at }}
                </a>
              </template>
            </b-table-column>

            <b-table-column v-slot="props" field="name" label="链接(link)" sortable>
              <template v-if="showDetailIcon">
                {{ props.row.link }}
              </template>
              <template v-else>
                <a @click="props.toggleDetails(props.row)">
                  {{ props.row.link }}
                </a>
              </template>
            </b-table-column>

            <template slot="detail" slot-scope="props">
              <article class="media">
                <figure class="media-left">
                  <!--   <p class="image is-64x64">
                    <img src="/static/img/placeholder-128x128.png">
                  </p>-->
                </figure>
                <div class="media-content">
                  <div class="content">
                    <p>
                      <strong>{{ props.row.name }} {{ props.row.name }}</strong>
                      <small>@{{ props.row.name }}</small>
                      <small>31m</small>
                      <br>
                      <strong>{{ props.row.title }} -- {{ props.row.sub_title }}</strong>
                    </p>
                  </div>
                </div>
              </article>
            </template>

          </b-table>
        </div>

      </section>
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

<script>
import { getBlogs } from '@/api/resources-blogs'
import { authuserGet } from '@/api/basic_user'
import { createArticle, fetchPv, updateArticle } from '@/api/article'
import { parseTime } from '@/utils'

// const data = require('@/data/sample.json')

export default {
  data() {
    return {
      // data,
      blogs: [],
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
  created() {
    // this.getList()
    this.resourceBlogsGet()
  },
  methods: {
    resourceBlogsGet() {
      this.listLoading = true
      getBlogs(this.listQuery).then(response => {
        //        this.list = response.data.items
        this.blogs = response.spec
        console.log('===>', this.blogs)
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
<style type="text/css">

</style>
