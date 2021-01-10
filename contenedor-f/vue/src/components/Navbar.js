import Vue from 'vue'

Vue.component('navbar',{
  template: 
  `
  <div>
    <nav>
      <div class="nav-wrapper teal darken-3">
        <div class="row" style="display: flex; align-items: center;">
          <button  data-target="slide-out" id="sidebarCollapse" style="color: white; display: flex; align-items: center;" class="waves-effect waves-teal btn-flat sidenav-trigger"><i class="material-icons">menu</i></button>
          <a href="#">{{page}}</a>
          <img class="brand-logo right" src="https://upload.wikimedia.org/wikipedia/commons/thumb/9/95/Vue.js_Logo_2.svg/1184px-Vue.js_Logo_2.svg.png" style="height: 35px;"/>
        </div>
      </div>
    </nav>
    <sidenav></sidenav>
  </div>
  `,
  props: ['page']
})