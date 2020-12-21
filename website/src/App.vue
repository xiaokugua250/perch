<template>

  <div id="app">
      <head>
             <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
      <link href="https://fonts.googleapis.com/css?family=Open+Sans" rel="stylesheet">
      <!-- Bulma Version 0.9.0-->
      <link rel="stylesheet" href="https://unpkg.com/bulma@0.9.0/css/bulma.min.css">
      <link rel="stylesheet" type="text/css" href="../css/hero.css">
      <link rel="stylesheet" href="https://unpkg.com/bulma-modal-fx/dist/css/modal-fx.min.css">
      </head>
    <router-view />
    <script type="application/ld+json">
      {
      "@context": "https://schema.org/",
      "@type": "TouristAttraction",
      "name": "z-gour.com",
      "description": "z-gour.com is a website Currently do not know what to do with the site,how about funny sites collections or useful tools to live;[还不知道用来干啥～]",
      "url": "https://z-gour.com",
      "isAccessibleForFree": true,
      "mainEntityOfPage": {
      "@type": "WebPage",
      "@id": "https://z-gour.com"
      }

      }
    </script>
  </div>

</template>

<script>
export default {
  name: 'App'
}

</script>
