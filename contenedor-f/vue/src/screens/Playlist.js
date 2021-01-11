import Vue from 'vue'

Vue.component('playlist',{
  template: 
  `
    <div>
      <navbar page="Playlist"></navbar>
      <div class="container">
        <div class="row valign-wrapper" style="height: 90vh;">
          <div class="carousel">
            <div v-for="song of songs" class="carousel-item" href="#one!">
              <media-player :song="song"></media-player>
            </div>
          </div>
        </div>
      </div>
    </div>
  `,
  data(){
    return {
      songs: ['Eve1', 'Eve2', 'Eve3','Eve1', 'Eve2', 'Eve3','Eve1', 'Eve2', 'Eve3']
    }
  }
})