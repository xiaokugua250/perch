<template>
  <div>

    <head>
      <meta charset="UTF-8">
      <title>Cheatsheet Concept</title>

      <!-- Bulma Version 0.9.0-->
      <link rel='stylesheet' href='https://unpkg.com/bulma@0.9.0/css/bulma.min.css'>
      <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@mdi/font@5.8.55/css/materialdesignicons.min.css">
      <link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.2.0/css/all.css">
    </head>
    <body>
    <router-view></router-view>
    <section class="hero is-primary">

      <div class="hero-body">
        <div style="margin-left: 200px;margin-right: 200px">
          <div class="container">
            <section class="section">
              <div class="container">
                <h1 class="title">平台应用</h1>
                <h2 class="subtitle">
                  A simple container to divide your page into <strong>sections</strong>, like the one you're currently reading
                </h2>
            <div class="columns" v-for="item in 2" :key="item">
              <div class="column"  v-for="item in 5" :key="item">
                <div class="card">
                  <div class="card-image">
                    <figure class="image ">
                      <img src="../../../../assets/images/sites/backend/monitor.png" alt="Placeholder image">
                    </figure>
                  </div>
                  <div class="card-content">
                  <!--  <div class="media">
                   <div class="media-left">
                        <figure class="image is-48x48">
                          <img src="../../../../assets/images/sites/backend/monitor.png" alt="Placeholder image">
                        </figure>
                      </div>
                      <div class="media-content">
                        <p class="title is-4">John Smith</p>
                        <p class="subtitle is-6">@johnsmith</p>
                      </div>
                    </div>-->
                    <div class="media-content">
                      <b-taglist>
                        <b-tag type="is-info">First</b-tag>
                        <b-tag type="is-info">Second</b-tag>
                        <b-tag type="is-info">Third</b-tag>

                      </b-taglist>
                    </div>
                    <div class="content">

                      Phasellus nec iaculis mauris. <a>@bulmaio</a>.
                      <a href="#">#css</a> <a href="#">#responsive</a>
                      <br>
                      <time datetime="2016-1-1">11:09 PM - 1 Jan 2016</time>
                      <b-button type="is-success is-light "   @click="isCardModalActive = true" >提交任务(submit)</b-button>

                    </div>



                  </div>
                </div>
              </div>
            </div>
              </div>
            </section>
          </div>

        </div>



      </div>
    </section>
    <section class="section">
      <div class="container">
        <div class="columns">
          <div class="column is-3">
            <aside class="is-medium menu">
              <b-menu>
                <b-menu-list label="Menu">
                  <b-menu-item icon="information-outline" label="Info"></b-menu-item>
                  <b-menu-item icon="settings" :active="isActive" expanded>
                    <template #label="props">
                      Administrator
                      <b-icon class="is-pulled-right" :icon="props.expanded ? 'menu-down' : 'menu-up'"></b-icon>
                    </template>
                    <b-menu-item icon="account" label="Users"></b-menu-item>
                    <b-menu-item icon="cellphone-link">
                      <template #label>
                        Devices
                        <b-dropdown aria-role="list" class="is-pulled-right" position="is-bottom-left">
                          <template #trigger>
                            <b-icon icon="dots-vertical"></b-icon>
                          </template>
                          <b-dropdown-item aria-role="listitem">Action</b-dropdown-item>
                          <b-dropdown-item aria-role="listitem">Another action</b-dropdown-item>
                          <b-dropdown-item aria-role="listitem">Something else</b-dropdown-item>
                        </b-dropdown>
                      </template>
                    </b-menu-item>
                    <b-menu-item icon="cash-multiple" label="Payments" disabled></b-menu-item>
                  </b-menu-item>
                  <b-menu-item icon="account" label="My Account">
                    <b-menu-item label="Account data"></b-menu-item>
                    <b-menu-item label="Addresses"></b-menu-item>
                  </b-menu-item>
                </b-menu-list>
                <b-menu-list>
                  <b-menu-item label="Expo" icon="link" tag="router-link" target="_blank" to="/expo"></b-menu-item>
                </b-menu-list>
                <b-menu-list label="Actions">
                  <b-menu-item label="Logout"></b-menu-item>
                </b-menu-list>
              </b-menu>
              <p class="menu-label">
                More to read...
              </p>
              <section>
                <b-field>
                  <b-tag v-if="isTag1Active"
                         type="is-primary"
                         closable
                         aria-close-label="Close tag"
                         @close="isTag1Active = false">
                    Colored closable tag label
                  </b-tag>
                </b-field>
                <b-field>
                  <b-tag v-if="isTag2Active"
                         type="is-success"
                         closable
                         closeType='is-danger'
                         aria-close-label="Close tag"
                         @close="isTag2Active = false">
                    Colored closable tag label win
                  </b-tag>
                </b-field>

                <b-field>
                  <b-tag v-if="isTag3Active"
                         attached
                         closable
                         aria-close-label="Close tag"
                         @close="isTag3Active = false">
                    Attached closable tag label
                  </b-tag>
                </b-field>

                <b-field>
                  <b-tag v-if="isTag4Active"
                         type="is-danger"
                         attached
                         closable
                         aria-close-label="Close tag"
                         @close="isTag4Active = false">
                    Colored attached closable tag label
                  </b-tag>
                </b-field>
                <b-field>
                  <b-tag v-if="isTag5Active"
                         close-type='is-danger'
                         attached
                         closable
                         aria-close-label="Close tag"
                         @close="isTag5Active = false">
                    Attached tag label with colored close type
                  </b-tag>
                </b-field>
                <b-field>
                  <b-tag v-if="isTag6Active"
                         close-type='is-warning'
                         close-icon-type='is-dark'
                         attached
                         closable
                         close-icon='delete'
                         aria-close-label="Close tag"
                         @close="isTag6Active = false">
                    Attached tag label with custom and colored icon
                  </b-tag>
                </b-field>
              </section>
            </aside>
          </div>
          <div class="column is-9">
            <div class="content is-medium">
              <h3 class="title is-3">Snippets ¯\_(ツ)_/¯</h3>
              <div class="box">
                <h4 id="basic_user" class="title is-3">平台软件(SoftWares)</h4>
                <article class="message is-primary">
                  <span class="icon has-text-primary">
                  <i class="fab fa-js"></i>
                  </span>
                  <div class="message-body">
                平台软件基本信息
                  </div>
                </article>
                <section>
                  <b-table
                    :data="data"
                    :paginated="isPaginated"
                    :per-page="perPage"
                    :detailed="true"
                    :detail-key="id"
                    :current-page.sync="currentPage"
                    :pagination-simple="isPaginationSimple"
                    :pagination-position="paginationPosition"
                    :default-sort-direction="defaultSortDirection"
                    :pagination-rounded="isPaginationRounded"
                    :sort-icon="sortIcon"
                    :sort-icon-size="sortIconSize"
                    default-sort="user.first_name"
                    aria-next-label="Next page"
                    aria-previous-label="Previous page"
                    aria-page-label="Page"
                    aria-current-label="Current page">

                    <b-table-column field="id" label="ID" width="40" sortable numeric v-slot="props">
                      {{ props.row.id }}
                    </b-table-column>

                    <b-table-column field="user.first_name" label="First Name" sortable v-slot="props">
                      {{ props.row.user.first_name }}
                    </b-table-column>

                    <b-table-column field="user.last_name" label="Last Name" sortable v-slot="props">
                      {{ props.row.user.last_name }}
                    </b-table-column>

                    <b-table-column field="date" label="Date" sortable centered v-slot="props">
                <span class="tag is-success">
{{ new Date(props.row.date).toLocaleDateString() }}
                </span>
                    </b-table-column>

                    <b-table-column label="Gender" v-slot="props">
                <span>
<b-icon pack="fas"
        :icon="props.row.gender === 'Male' ? 'mars' : 'venus'">
</b-icon>
{{ props.row.gender }}
                </span>
                    </b-table-column>
                    <template #detail="props">
                      <article class="media">
                        <figure class="media-left">
                          <p class="image is-64x64">
                            <img src="/static/img/placeholder-128x128.png">
                          </p>
                        </figure>
                        <div class="media-content">
                          <div class="content">
                            <p>
                              <strong>{{ props.row.user.first_name }} {{ props.row.user.last_name }}</strong>
                              <small>@{{ props.row.user.first_name }}</small>
                              <small>31m</small>
                              <br>
                              Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                              Proin ornare magna eros, eu pellentesque tortor vestibulum ut.
                              Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis.
                            </p>
                          </div>
                        </div>
                      </article>
                    </template>
                  </b-table>
                </section>


              </div>

              <div class="box">
                <h4 id="basic_permission" class="title is-3">应用实例(ApplicationInstances)</h4>
                <article class="message is-primary">
                  <span class="icon has-text-primary">
                  <i class="fab fa-js"></i>
                  </span>
                  <div class="message-body">
                    应用实例基本信息
                  </div>
                </article>
                <section>
                  <b-table
                    :data="data"
                    :paginated="isPaginated"
                    :per-page="perPage"
                    :detailed="true"
                    :detail-key="id"
                    :current-page.sync="currentPage"
                    :pagination-simple="isPaginationSimple"
                    :pagination-position="paginationPosition"
                    :default-sort-direction="defaultSortDirection"
                    :pagination-rounded="isPaginationRounded"
                    :sort-icon="sortIcon"
                    :sort-icon-size="sortIconSize"
                    default-sort="user.first_name"
                    aria-next-label="Next page"
                    aria-previous-label="Previous page"
                    aria-page-label="Page"
                    aria-current-label="Current page">

                    <b-table-column field="id" label="ID" width="40" sortable numeric v-slot="props">
                      {{ props.row.id }}
                    </b-table-column>

                    <b-table-column field="user.first_name" label="First Name" sortable v-slot="props">
                      {{ props.row.user.first_name }}
                    </b-table-column>

                    <b-table-column field="user.last_name" label="Last Name" sortable v-slot="props">
                      {{ props.row.user.last_name }}
                    </b-table-column>

                    <b-table-column field="date" label="Date" sortable centered v-slot="props">
                <span class="tag is-success">
{{ new Date(props.row.date).toLocaleDateString() }}
                </span>
                    </b-table-column>

                    <b-table-column label="Gender" v-slot="props">
                <span>
<b-icon pack="fas"
        :icon="props.row.gender === 'Male' ? 'mars' : 'venus'">
</b-icon>
{{ props.row.gender }}
                </span>
                    </b-table-column>
                    <template #detail="props">
                      <article class="media">
                        <figure class="media-left">
                          <p class="image is-64x64">
                            <img src="/static/img/placeholder-128x128.png">
                          </p>
                        </figure>
                        <div class="media-content">
                          <div class="content">
                            <p>
                              <strong>{{ props.row.user.first_name }} {{ props.row.user.last_name }}</strong>
                              <small>@{{ props.row.user.first_name }}</small>
                              <small>31m</small>
                              <br>
                              Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                              Proin ornare magna eros, eu pellentesque tortor vestibulum ut.
                              Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis.
                            </p>
                          </div>
                        </div>
                      </article>
                    </template>
                  </b-table>
                </section>


              </div>
              <div class="box">
                <h4 id="const" class="title is-3">const</h4>
                <article class="message is-primary">
                  <span class="icon has-text-primary">
                  <i class="fab fa-js"></i>
                  </span>
                  <div class="message-body">
                    Block-scoped. Cannot be re-assigned. Not immutable.
                  </div>
                </article>
                <pre><code class="language-javascript">const test = 'test';</code></pre>
              </div>
              <div class="box">
                <h4 id="let" class="title is-3">let</h4>
                <article class="message is-primary">
                <span class="icon has-text-primary">
                  <i class="fas fa-info-circle"></i>
                </span>
                  <div class="message-body">
                    Block-scoped. Can be re-assigned.
                  </div>
                </article>
                <pre><code class="language-javascript">let i = 0;</code></pre>
              </div>
              <h3 class="title is-3">More to Come...</h3>
              <div class="box">
                <h4 id="lorem" class="title is-4">More to come...</h4>
                <article class="message is-primary">
              <span class="icon has-text-primary">
                <i class="fas fa-info-circle"></i>
              </span>
                  <div class="message-body">
                    Lorem ipsum dolor sit amet, mea ne viderer veritus menandri, id scaevola gloriatur instructior sit.
                  </div>
                </article>
                <pre><code class="language-javascript">let i = 0;</code></pre>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>


    </body>

  </div>

</template>


<script>
export default {
  data() {
    const data = [{"id":1,"user":{"first_name":"Jesse","last_name":"Simmons"},"date":"2016/10/15 13:43:27","gender":"Male"},
      {"id":2,"user":{"first_name":"John","last_name":"Jacobs"},"date":"2016/12/15 06:00:53","gender":"Male"},
      {"id":3,"user":{"first_name":"Tina","last_name":"Gilbert"},"date":"2016/04/26 06:26:28","gender":"Female"},
      {"id":4,"user":{"first_name":"Clarence","last_name":"Flores"},"date":"2016/04/10 10:28:46","gender":"Male"},
      {"id":5,"user":{"first_name":"Anne","last_name":"Lee"},"date":"2016/12/06 14:38:38","gender":"Female"},
      {"id":6,"user":{"first_name":"Sara","last_name":"Armstrong"},"date":"2016/09/23 18:50:04","gender":"Female"},
      {"id":7,"user":{"first_name":"Anthony","last_name":"Webb"},"date":"2016/08/30 23:49:38","gender":"Male"},
      {"id":8,"user":{"first_name":"Andrew","last_name":"Greene"},"date":"2016/11/20 14:57:47","gender":"Male"},
      {"id":9,"user":{"first_name":"Russell","last_name":"White"},"date":"2016/07/13 09:29:49","gender":"Male"},

      {"id":60,"user":{"first_name":"Christopher","last_name":"Palmer"},"date":"2016/05/24 08:58:12","gender":"Male"}]

    return {
      isTag1Active:true,
      isTag2Active:true,
      isTag3Active:true,
      isTag4Active:true,
      isTag5Active:true,
      isTag6Active:true,
      isActive:true,
      isCardModalActive: false,
      data,

      isPaginated: true,
      isPaginationSimple: true,
      isPaginationRounded: false,
      paginationPosition: 'bottom',
      defaultSortDirection: 'asc',
      sortIcon: 'arrow-up',
      sortIconSize: 'is-small',
      currentPage: 1,
      perPage: 5
    }
  }
}


/*document.addEventListener("DOMContentLoaded", function() {
  //The first argument are the elements to which the plugin shall be initialized
  //The second argument has to be at least a empty object or a object with your desired options
  OverlayScrollbars(document.querySelectorAll("body"), { });
});*/

</script>
