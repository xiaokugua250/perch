<template>
  <div id="app">

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
