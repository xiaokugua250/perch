<template>
  <div>

    <head>
      <meta charset="utf-8">
      <meta http-equiv="X-UA-Compatible" content="IE=edge">
      <meta name="viewport" content="width=device-width, initial-scale=1">
      <meta name="google-site-verification" content="oRlwQCWDSEBrMAdKd3t3_3pCA9sFayLgZdgLer2TKfM">
      <link rel="canonical" href="https://booknshelf.com">
      <meta name="keywords" content="booknshelf, catalog your books, online free library, catalogue your books, book cataloging, free book catalog, catalogue, library books online">

      <meta
        name="description"
        content="Your reading journey in a single place."
      >
      <meta name="google-site-verification" content="oRlwQCWDSEBrMAdKd3t3_3pCA9sFayLgZdgLer2TKfM">
      <meta name="google-site-verification" content="WrDufWmHHT4G7Vp90CsujULhVdDrlmb48mfZZUoroN4">

      <meta property="fb:app_id" content="1899203000306326">
      <meta property="og:title" content="Keep track of all your books online">
      <meta
        property="og:description"
        content="Your reading journey in a single place."
      >
      <meta property="og:image" content="https://booknshelf.com/img/social/homepage-screenshot-social-main-new.png">
      <meta property="og:image:height" content="717">
      <meta property="og:image:width" content="1200">
      <meta property="og:site_name" content="Booknshelf">
      <meta property="og:type" content="website">
      <meta property="og:url" content="https://booknshelf.com">

      <meta name="twitter:card" content="summary">
      <meta
        name="twitter:description"
        content="Stay organized with your books."
      >
      <meta
        name="twitter:image"
        content="https://booknshelf.com/img/social/homepage-screenshot-social-main-new.png"
      >
      <meta name="twitter:site" content="@booknshelf">
      <meta
        name="twitter:title"
        content="Catalog your books online. Without hustle."
      >

      <!-- CSRF Token -->
      <meta name="csrf-token" content="mRNRamrw8XWWc9Sb6bapaALTt7iXeNacfU0z3Pib">

      <title>Booknshelf: Keep track of all your books online.</title>
      <link rel="icon" type="image/png" href="/img/favicons/favicon-16x16.png" sizes="16x16">
      <link rel="icon" type="image/png" href="/img/favicons/favicon-32x32.png" sizes="32x32">
      <link rel="icon" type="image/png" href="/img/favicons/favicon-96x96.png" sizes="96x96">

      <!-- Styles -->
      <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/hover.css/2.1.0/css/hover-min.css">
      <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.6.3/css/font-awesome.min.css">
      <link rel="stylesheet" href="/build/css/booknshelf-2ed947b2c2.css">

    </head>
    <body>
      <div id="app">
        <!-- Navigation -->
        <!-- NavBar For Guest Users -->
        <nav class="navbar" role="navigation" aria-label="main navigation">
          <div class="container">
            <div class="navbar-brand">
              <a class="navbar-item" href="/">
                <strong>🏠 Home</strong>
              </a>
              <button class="button navbar-burger" data-target="navMenu">
                <span />
                <span />
                <span />
              </button>
            </div>
            <div id="navMenu" class="navbar-menu">
              <div class="navbar-start">
                <!-- navbar items -->
                <a class="navbar-item" href="/faq">
                  <span>FAQ</span>
                </a>
                <a class="navbar-item" href="/blog">
                  <span>Blog</span>
                </a>
                <a class="navbar-item" href="/topics">
                  Topics
                </a>
                <a class="navbar-item" href="/bookshelves">
                  Bookshelves
                </a>
                <a class="navbar-item is-warning" href="https://paypal.me/tiggreen" target="_blank">
                  <span class="icon">
                    <i class="fa fa-paypal" />
                  </span>
                  <span><strong>Make a donation</strong></span>
                </a>
              </div>
              <div class="navbar-end">
                <!-- navbar items -->
                <div class="changelog navbar-item is-hidden-touch">
                  <a href="#" class="changloglink" style="color: #08ac94">What's New</a>
                </div>
                <div class="navbar-item">
                  <a class="button is-primary" href="/register">SIGN UP</a>
                </div>
                <div class="navbar-item">
                  <a class="button is-light" href="/login">LOG IN</a>
                </div>
              </div>
            </div>
          </div>
        </nav>

        <!-- Application Level Modals
      <please-login-modal v-if="plaseLoginModal" @close="plaseLoginModal = false"></please-login-modal>
-->
        <!-- Main Content -->
        <section class="hero is-medium home-search-bg" style="background-image: url('../../assets/images/sites/search/search_background.jpg')">
          <div class="hero-body home-search-body">
            <div class="container has-text-centered">
              <div class="columns">
                <div class="column" />
                <div class="column is-6">
                  <h3 class="title is-3 primary-span-home">
                    Keep track of all your books online. Without hassle.
                  </h3>
                  <p /><form role="form" method="GET" action="https://booknshelf.com/books/search">
                    <p class="control has-icon">
                      <input
                        class="input is-large"
                        type="text"
                        value=""
                        name="q"
                        placeholder="Search for great books here ..."
                      >
                      <span class="icon">
                        <i class="fa fa-search" />
                      </span>
                    </p>
                    <input type="submit" style="display: none;">
                  </form>
                  </p>
                </div>
                <div class="column" />
              </div>
            </div>
          </div>
        </section>

        <section class="section" style="padding-bottom: 5px;">
          <div class="container">
            <div class="columns">
              <div class="column is-8 is-offset-2">
                <article class="message is-primary" style="background-color: transparent;">
                  <div class="message-body">
                    <p class="title has-text-centered">
                      A few things Booknshelf can help you with
                    </p>
                  </div>
                </article>
              </div>
            </div>
          </div>
        </section>

        <section class="section">
          <div class="container">
            <div class="columns">
              <div class="column is-3 has-text-centered">
                <svg xmlns="http://www.w3.org/2000/svg" width="100" height="100" viewBox="0 0 100 100" class="heroicon-magnify heroicon heroicons-lg">
                  <path class="heroicon-magnify-glass-edge heroicon-component-accent heroicon-component-fill" fill="#7ACFC0" d="M70 36a34 34 0 1 1-68 0 34 34 0 0 1 68 0z" />
                  <path class="heroicon-magnify-glass heroicon-component-fill" fill="#FFFFFF" d="M61 36a25 25 0 1 1-50 0 25 25 0 0 1 50 0z" />
                  <polygon class="heroicon-magnify-handle heroicon-component-fill" fill="#FFFFFF" points="94.879 87.707 75.293 68.121 68.121 75.293 87.707 94.879" />
                  <path class="heroicon-magnify-handle-connector heroicon-component-fill" fill="#FFFFFF" d="M63.92 58.73L65.16 60 60 65.17l-1.26-1.26a36.22 36.22 0 0 0 5.18-5.18zm-.5 5.86l1.17-1.18 4 4-1.18 1.18-4-4z" />
                  <path class="heroicon-magnify-handle-edge heroicon-component-accent heroicon-component-fill" fill="#7ACFC0" d="M90 97.17l-1.59-1.58 7.18-7.18L97.17 90 90 97.17zM73 65.83l1.59 1.58-7.18 7.18L65.83 73 73 65.83z" />
                  <path class="heroicon-shadows" fill="#000000" d="M6.04 37.5a30 30 0 1 1 59.93 0 30 30 0 0 0-59.93 0z" opacity=".2" />
                  <path class="heroicon-outline" fill="#4A4A4A" fill-rule="nonzero" d="M65.14 57.14l1.45 1.45L68 60l-1.41 1.41L66 62l4 4 1.59-1.59L73 63l1.41 1.41L76 66l.7.7 19.6 19.6.7.7 1.59 1.59L100 90l-1.41 1.41-7.18 7.18L90 100l-1.41-1.41L87 97l-.7-.7-19.6-19.6-.7-.7-1.59-1.59L63 73l1.41-1.41L66 70l-4-4-.59.59L60 68l-1.41-1.41-1.45-1.45h.01a36 36 0 1 1 8-8zM70 36a34 34 0 1 0-68 0 34 34 0 0 0 68 0zm-6.08 22.73a36.22 36.22 0 0 1-5.18 5.18L60 65.17 65.17 60l-1.26-1.26zm-.5 5.86l4 4 1.17-1.18-4-4-1.18 1.18zM90 97.17L97.17 90l-1.58-1.59-7.18 7.18L90 97.17zm4.88-9.46L75.29 68.12l-7.17 7.17 19.59 19.59 7.17-7.17zM73 65.83L65.83 73l1.58 1.59 7.18-7.18L73 65.83zM66 36a30 30 0 1 1-60 0 30 30 0 0 1 60 0zM36 65a29 29 0 1 0 0-58 29 29 0 0 0 0 58zm0-2a27 27 0 1 1 0-54 27 27 0 0 1 0 54zm25-27a25 25 0 1 0-50 0 25 25 0 0 0 50 0zM36 15v1a20 20 0 0 0-17.9 28.95l-.89.44A21 21 0 0 1 36 15zm37.65 61.35l.7-.7 14 14-.7.7-14-14z" />
                </svg>

                <p class="subtitle has-text-centered" style="margin-top: 6px;">
                  <strong>Search for books or discover from our curated topics</strong>
                </p>
                <p class="subtitle has-text-centered" style="margin-top: -0.5rem;">
                  We're using Amazon's huge book database for search.
                </p>
              </div>
              <div class="column is-3 has-text-centered">
                <svg xmlns="http://www.w3.org/2000/svg" width="100" height="100" viewBox="0 0 100 100" class="heroicon-book heroicon heroicons-lg">
                  <path class="heroicon-book-cover heroicon-component-fill" fill="#FFFFFF" d="M89 75H15v10h10v4h-9a5 5 0 0 1-5-5V6a5 5 0 0 1 5-5h73v74zM39 85h50v4H39v-4z" />
                  <path class="heroicon-book-binding heroicon-component-accent heroicon-component-fill" fill="#7ACFC0" d="M25 71V1h-9a5 5 0 0 0-5 5v69l4-4h10z" />
                  <rect class="heroicon-book-pages heroicon-component-fill" width="72" height="10" x="15" y="75" fill="#FFFFFF" />
                  <polygon class="heroicon-book-bookmark heroicon-component-accent heroicon-component-fill" fill="#7ACFC0" points="24 77 40 77 40 99 32 95 24 99" />
                  <path class="heroicon-shadows" fill="#000000" d="M11 75.99A4.99 4.99 0 0 1 16 71h73v4h-2v10h2v4H39v-7H25v7h-9a5 5 0 0 1-5-4.99v-8.02z" opacity=".2" />
                  <path class="heroicon-outline" fill="#4A4A4A" fill-rule="nonzero" d="M24 90h-8a6 6 0 0 1-6-6V6A6 6 0 0 1 16 0h74v76h-2v8h2v6H40v10l-8-4-8 4V90zM88 2H20v66h-1V2h-2v66h-1V2a4 4 0 0 0-4 4.01v65.52A5.97 5.97 0 0 1 16 70h72V2zM16 72a4 4 0 0 0-4 4V84A4 4 0 0 0 16 88h8v-2h-8a2 2 0 0 1-2-2v-8c0-1.1.9-2 2-2h72v-2H16zm24 16h48v-2H40v2zm0-4h46v-1H40v1zm0-2h46v-1H58v-1h28v-2H40v4zm-2-4H26v18.76l6-3 6 3V78zm-14 0h-6v-1h68v-1H16v8h8v-6zm0-76h2v68h-2V2zm18 18h31v2H41v-2h1zm3 5h24v2H45v-2z" />
                </svg>
                <p class="subtitle has-text-centered" style="margin-top: 6px;">
                  <strong>Easily catalog all your books in shelves</strong>
                </p>
                <p class="subtitle has-text-centered" style="margin-top: -0.5rem;">
                  Never forget any book you've read or want to read.
                </p>
              </div>
            <!---
            参考：https://booknshelf.com/ 对页面进行补全
    -->
            </div>
          </div>
        </section>

        <section class="section" style="padding-bottom: 5px;">
          <div class="container">
            <div class="columns">
              <div class="column is-8 is-offset-2">
                <article class="message is-primary" style="background-color: transparent;">
                  <div class="message-body">
                    <p class="title has-text-centered">
                      Check out Booknshelf's profile for inspiration
                    </p>
                  </div>
                </article>
              </div>
            </div>
          </div>
        </section>

        <section class="section">
          <div class="container">
            <div class="columns">
              <div class="column is-8 is-offset-2">
                <div class="card">
                  <div class="card-image">
                    <figure class="image profile-figure">
                      <img
                        src="https://booknshelf.imgix.net/homepage/booknshelf-profile-showcase.png"
                        alt="Booknshelf Profile Homepage Showcase"
                        class="profile-figure-image"
                      >
                      <div class="middle">
                        <a class="button is-primary is-large" target="_blank" href="/@booknshelf">
                          <span class="icon">
                            <i class="fa fa-external-link" />
                          </span>
                          <span> Go to profile</span>
                        </a>
                      </div>

                    </figure>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>

        <section class="hero">
          <div class="hero-body">
            <div class="container">
              <h1 class="title has-text-centered">
                <a class="button is-primary is-large hvr-float-shadow" href="https://booknshelf.com/register">
                  <strong>Get started now</strong>
                </a>
              </h1>
            </div>
          </div>
        </section>

        <section class="section is-primary is-fullheight is-bold">
          <div class="container">
            <div class="notification has-text-centered">
              <div class="title is-4">
                <strong>Discover new books from some of our favorite topics below </strong>
                <a href="/topics" style="color: #00AB91;">See all ></a>
              </div>
            </div>
          </div>
        </section>

        <section class="section is-primary is-fullheight is-bold">
          <div class="container">
            <div class="notification has-text-centered">
              <div class="title is-4">
                <strong>Check out some of our popular books below</strong>
              </div>
            </div>
          </div>
        </section>

        <footer class="footer">
          <div class="container">
            <div class="content has-text-centered">
              <p>
                <strong>Booknshelf</strong> by <a
                  style="text-decoration: underline;"
                  href="https://tigran.nyc"
                  target="_blank"
                >Tigran</a>
              </p>
              <div id="social">
                <iframe
                  id="twitter-widget-0"
                  scrolling="no"
                  frameborder="0"
                  allowtransparency="true"
                  class="twitter-follow-button twitter-follow-button-rendered"
                  title="Twitter Follow Button"
                  src="//platform.twitter.com/widgets/follow_button.c4fd2bd4aa9a68a5c8431a3d60ef56ae.en.html#dnt=false&amp;id=twitter-widget-0&amp;lang=en&amp;screen_name=tiggreen&amp;show_count=true&amp;show_screen_name=true&amp;size=m&amp;time=1485106087141"
                  style="position: static; visibility: visible; width: 203px; height: 20px;"
                  data-screen-name="tiggreen"
                />
                <iframe
                  src="https://www.facebook.com/plugins/like.php?href=https%3A%2F%2Fwww.facebook.com%2Fusebooknshelf%2F&width=61&layout=button_count&action=like&size=small&show_faces=false&share=false&height=21&appId=1899203000306326"
                  width="70"
                  height="21"
                  style="border:none;overflow:hidden"
                  scrolling="no"
                  frameborder="0"
                  allowTransparency="true"
                />
                <iframe
                  id="twitter-widget-1"
                  scrolling="no"
                  frameborder="0"
                  allowtransparency="true"
                  class="twitter-share-button twitter-share-button-rendered twitter-tweet-button"
                  title="Twitter Tweet Button"
                  src="//platform.twitter.com/widgets/tweet_button.c4fd2bd4aa9a68a5c8431a3d60ef56ae.en.html#dnt=false&amp;id=twitter-widget-1&amp;lang=en&amp;original_referer=https://booknshelf.com&amp;related=tiggreen%3ACreator%20of%20Booknshelf&amp;size=m&amp;text=Keep track of all your books online!&amp;time=1485106087144&amp;type=share&amp;url=https%3A%2F%2Fbooknshelf.com&amp;via=booknshelf"
                  style="position: static; visibility: visible; width: 61px; height: 20px;"
                  data-url="https://booknshelf.com"
                />
              </div>
              <p>Consider a small donation to help me to keep the site running!</p>
              <a class="button is-dark btn btn-bright" href="https://paypal.me/tiggreen" target="_blank" type="button">
                <span class="icon">
                  <i class="fa fa-paypal" />
                </span>
                <span>MAKE A DONATION!</span>
              </a>
              </p>
            </div>
          </div>
        </footer>

      </div>

    <!-- Scripts -->

    </body>
  </div>
</template>
 <!-- <script src="../js/bulma.js"></script> -->
