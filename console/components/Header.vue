<template>
  <header class="header fn-flex">
    <div class="header__logo" v-if="$route.path.indexOf('admin') > -1">
      <img src="~static/images/logo.png"/>
      Solo
    </div>
    <div class="header__nav fn-flex-1">
      user logout
    </div>
  </header>
</template>

<script>
  export default {
    props: ['isLogin']
  }
</script>

<style lang="sass">
  @import '~assets/scss/_variables.scss'

  .header
    position: fixed
    height: 60px
    width: 100%
    z-index: 10
    top: 0

    &__logo
      background-color: $white
      width: 240px

    &__nav
      background-color: $blue
</style>