<template>
  <div>
    <div class="sp-checkers__main sp-box sp-shadow sp-form-group">
        <form class="sp-checkers__main__form">
          <div class="sp-checkers__main__rcpt__header sp-box-header">
            Create a Checkers Game
          </div>

          <input class="sp-input" placeholder="Red: cosmos1xxx" v-model="red" />
          <input class="sp-input" placeholder="Black: cosmos1xxx" v-model="black" />
          <sp-button @click="submit">Create game</sp-button>
        </form>
    </div>
  </div>
</template>
<script>
export default {
  name: "CreateGame",
  data() {
    return {
      red: "",
      black: "",
    };
  },
  computed: {

    currentAccount() {
      if (this._depsLoaded) {
        if (this.loggedIn) {
          return this.$store.getters['common/wallet/address']
        } else {
          return null
        }
      } else {
        return null
      }
    },
    loggedIn() {
      if (this._depsLoaded) {
        return this.$store.getters['common/wallet/loggedIn']
      } else {
        return false
      }
    }
  },
  methods: {
    async submit() {
      const value = {
        creator: this.currentAccount,
        red: this.red,
        black: this.black,
      };
      const result = await this.$store.dispatch("b9lab.checkers.checkers/sendMsgCreateGame", {
        value,
        fee: [],
      });
      console.log(result);
    },
  },
};
</script>
