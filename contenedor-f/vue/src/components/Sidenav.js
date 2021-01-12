import Vue from 'vue'

Vue.component('sidenav',{
  template: 
  `
  <ul id="slide-out" class="sidenav">
    <li>
        <div class="user-view" style="padding-bottom: 100px;">
            <div class="background">
                <img style="height: 120px;" src="https://i.pinimg.com/originals/85/b0/db/85b0db4f7ad7dce7ff28e14f7db75e17.png">
            </div>
        </div>
    </li>
    </ul>
  `,
  props: ['page']
})
