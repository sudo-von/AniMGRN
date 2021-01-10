import Vue from 'vue'

Vue.component('reproductor',{
  template: 
  `
    <nav class="navbar navbar-expand-lg navbar-light" style="background: #198754">
      <button class="btn transparent text-white d-flex align-items-center"><i class="material-icons">menu</i></button>
        <a class="navbar-brand text-white" href="#">{{page}}</a>
    </nav>
  `,
  data(){
    return {
      page: 'Playlist'
    }
  }
})