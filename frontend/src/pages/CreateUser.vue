<script setup>
import BmButton from '../components/atoms/BmButton.vue'
import BmInput from '../components/atoms/BmInput.vue'
import { router } from '../router.js'
import { ref } from 'vue'

const fullName = ref('')
const email = ref('')
const password = ref('')
const showPassword = ref(false)

function createUser() {
  // Functionality to create a new user will be implemented here
    fetch('http://localhost:8000/api/user/', {
      method: "POST", // Or POST, PUT, DELETE, etc.
      body: JSON.stringify({
        name: fullName.value,
        email: email.value,
        password: password.value
      }),
      headers: {
          "Content-Type": "application/json" // Example for JSON data
      },
    })
    .then(response => response.json()) // Parse the response as JSON
    .then(data => {
      console.log(data); // Work with the fetched data
      router.push({ name: 'login' });
    })
    .catch(error => {
      console.error('Error fetching data:', error);
    });
}

function redirectLogin() {
  router.push({ name: 'login' });
}
</script>

<template>
<div class="bg-background-light dark:bg-[#121212] font-display text-[#EAEAEA]">
  <div class="relative flex min-h-screen w-full flex-col items-center justify-center overflow-hidden p-4">
    <div aria-hidden="true" class="absolute -top-1/4 -left-24 h-96 w-96 rounded-full bg-primary/20 blur-[120px] filter"></div>
    <div aria-hidden="true" class="absolute -bottom-1/4 -right-24 h-96 w-96 rounded-full bg-primary/10 blur-[100px] filter"></div>
    <main class="z-10 flex w-full max-w-md flex-col items-center">
      <div class="mb-8 flex flex-col items-center gap-3">
        <span class="material-symbols-outlined text-primary text-4xl">pest_control</span>
        <h1 class="text-4xl font-bold text-white">Bug Marker</h1>
      </div>
      <div class="w-full rounded-xl border border-white/10 bg-black/20 p-8 backdrop-blur-sm">
        <h2 class="mb-6 text-center text-3xl font-bold tracking-tight text-white">Crie sua Conta</h2>
        <form class="flex flex-col gap-5" @submit.prevent>
          <label class="flex flex-col">
            <p class="pb-2 text-base font-medium leading-normal text-[#EAEAEA]">Nome Completo</p>
            <BmInput
              type="text"
              placeholder="Insira seu nome completo"
              v-model="fullName"
            />
          </label>
          <label class="flex flex-col">
            <p class="pb-2 text-base font-medium leading-normal text-[#EAEAEA]">E-mail</p>
            <BmInput
              type="email"
              placeholder="seuemail@exemplo.com"
              v-model="email"
            />
          </label>
          <label class="flex flex-col">
            <p class="pb-2 text-base font-medium leading-normal text-[#EAEAEA]">Senha</p>
            <div class="flex w-full flex-1 items-stretch rounded-lg">
              <BmInput
                :type="showPassword ? 'text' : 'password'"
                placeholder="Crie uma senha segura"
                v-model="password"
                noRightBorder
              />
              <div class="text-subtle-text-light dark:text-subtle-text-dark flex border border-border-light dark:border-border-dark bg-white dark:bg-background-dark items-center justify-center pr-[15px] rounded-r-lg border-l-0 cursor-pointer"
                   @click="showPassword = !showPassword">
                <span class="material-symbols-outlined">{{ showPassword ? 'visibility_off' : 'visibility' }}</span>
              </div>
            </div>
          </label>
          <BmButton @click="createUser" type="submit" class="mt-4">Cadastrar</BmButton>
        </form>
      </div>
      <div class="mt-6 text-center">
        <p class="text-sm text-gray-400">
          Já tem uma conta? 
          <a @click="redirectLogin" class="font-bold text-primary hover:underline" href="#">Faça login</a>
        </p>
      </div>
    </main>
  </div>
</div>
</template>