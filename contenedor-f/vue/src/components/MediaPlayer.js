import Vue from 'vue'

Vue.component('media-player',{
  template: 
  `
    <div class="card z-depth-1">
      <div class="row">
        <div class="col s12">
          <div class="col s12 m3" style="display: flex; justify-content: center;">
            <div class="card-content" style="display: flex; justify-content: center;">
              <img class="z-depth-1" style="border-radius: 100px;" height="60" width="60" src="https://yt3.ggpht.com/ytc/AAUvwnglGbT5N8Jwfphy_wP2DWKXZ2mkbY3q1ZfmA44O1w=s900-c-k-c0x00ffffff-no-rj">
            </div>
          </div>
          <div class="col s12 m9">
            <ul>                
                <li style="font-size: 18px; font-weight: bold;">{{song}}</li>
                <li style="font-size: 14px; color: gray;">Eve - Namiyos</li>
                <li style="font-size: 11px; color: gray;">2020-02-02</li>
                <li style="font-size: 11px; color: gray;">4 Minutos</li>
              </ul>
          </div>
        </div>
        <div class="col s12" style="margin-bottom: 10px;">
          <div class="video-container">
            <iframe src="https://www.youtube.com/embed/BEEFXAltoqo" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
          </div>
        </div>
      </div>
    </div>
  `,
  data(){
    return {
      page: 'Playlist'
    }
  },
  props: ['song']
})