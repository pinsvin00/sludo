<template>
  <div class="main">
    <h1 id="main-caption">Welcome in SuperLudo!</h1>
    <div class="login-box">
      <md-field>
        <label>Type in your nick!</label>
        <md-input v-model="nickname"></md-input>
        <span class="md-helper-text">Please dont use swear words!</span>
      </md-field>
      <md-button @click="redirectToGameView">PLAY</md-button>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import * as api from "@/store/api";
import axios from "axios";
import Cookies from "js-cookie";
@Component({
  components: {},
})
export default class Home extends Vue {
  nickname = "";
  async redirectToGameView() {
    const response = await api.registerPlayer(this.nickname);
    if (response.data.success === "true") {
      Cookies.set("key", response.data.key);
      Cookies.set("player_name", this.nickname);
      Cookies.set("player_number", response.data.player_number);
      Cookies.set("game_id", response.data.game_id);
      console.log(response.data);
      this.$router.push("Game");
    }
  }
}
</script>

<style scoped>
.main {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
}
#main-caption {
  font-size: 72px;
  margin-bottom: 10rem;
}
.background {
  z-index: -1;
  position: absolute;
  width: 100%;
  height: 100%;
}
.login-box {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 10%;
  justify-content: center;
}
</style>
