<template>
  <v-app>
    <v-main style="margin: 1rem">
      <v-data-table
        :headers="headers"
        :items="users"
        :items-per-page="5"
        class="elevation-1"
      >
      </v-data-table>
    </v-main>
  </v-app>
</template>

<script>

export default {
  name: 'Home',
  data() {
    return {
      headers: [
        { text: '아이디', value: 'id' },
        { text: '이름', value: 'name' },
        { text: '이메일', value: 'email' },
      ],
      users: []
    }
  },
  methods: {
    async updateUserList() {
      try {
        const req = await fetch("/api/user/list");
        this.users = await req.json();
      } catch(e) {
        console.log("fail to get user list: ", e);
      }
    }
  },
  created() {
    this.updateUserList();
  },
  components: {
  }
}
</script>
