<template>
  <div>
    <!-- Title Page -->
    <section
      class="bg-title-page p-t-40 p-b-50 flex-col-c-m"
      :style="{
        backgroundImage:
          'url(' + require('@/assets/login-background.jpeg') + ')',
      }"
    >
      <h2 class="l-text2 t-center">Login</h2>
    </section>

    <!-- content page -->
    <section class="bgwhite p-t-66 p-b-60">
      <div class="container">
        <div class="login-form m-auto">
          <Form @submit.prevent="login" :validation-schema="schema">
            <p class="m-b-20 text-center text-danger">
              {{ loginMessage }}
            </p>

            <div class="bo4 of-hidden size15 m-b-10">
              <Field
                  v-model="user.email"
                name="email"
                type="text"
                placeholder="Email"
                class="sizefull s-text7 p-l-22 p-r-22"
                :disabled="isLoading"
              />
            </div>

            <ErrorMessage name="email" class="text-danger m-b-20 d-block" />

            <div class="bo4 of-hidden size15 m-b-10">
              <Field
                  v-model="user.password"
                name="password"
                type="password"
                placeholder="Password"
                class="sizefull s-text7 p-l-22 p-r-22"
                :disabled="isLoading"
              />
            </div>

            <ErrorMessage name="password" class="text-danger m-b-20 d-block" />

            <div class="w-size25 m-auto">
              <button
                  @click="login"
                class="
                  flex-c-m
                  size2
                  bg1
                  bo-rad-23
                  hov1
                  m-text3
                  trans-0-4
                  m-t-20
                "
                :class="{ disabled: isLoading }"
                :disabled="isLoading"
              >
                <span v-show="isLoading" data-loader="ball-scale"></span>
                Login
              </button>
            </div>
          </Form>

          <br />
          <hr />

          <p class="text-center">
            <router-link to="/register"
              >Don't have an account yet? Register now!
            </router-link>
          </p>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import {mapActions, mapState} from "vuex";
import { Form, Field, ErrorMessage } from "vee-validate";
import * as yup from "yup";

export default {
  name: "LoginPage",

  components: {
    Form,
    Field,
    ErrorMessage,
  },

  data() {
    const schema = yup.object().shape({
      email: yup
        .string()
        .required("Email is required!")
        .email("Email is invalid!"),
      password: yup.string().required("Password is required!"),
    });

    return {
      isLoading: false,
      schema,
      user: {
        email: '',
        password: ''
      }
    };
  },

  computed: mapState("users", ["isLoginSuccess", "loginMessage", "user"]),

  created() {
    if (this.isLoginSuccess) {
      console.log(this.isLoginSuccess)
      this.$router.push("/");
    }
  },

  methods: {
    ...mapActions("cart", ["getCartById"]),
    login() {
      this.isLoading = true;
      this.$store.dispatch("users/login", {
        email: this.user.email,
        password: this.user.password,
      })
          .then(() => {
            this.getCartById(this.$store.state.users.user.id)
      })
          .then(() => {
        this.isLoading = false;
        if (this.isLoginSuccess) {
          this.$router.push("/");
        }
      })
    },
  },
};
</script>

<style scoped>
.login-form {
  max-width: 500px;
}
button.disabled {
  opacity: 0.7;
  cursor: not-allowed;
}
button.disabled:hover {
  background-color: #222222;
}
</style>