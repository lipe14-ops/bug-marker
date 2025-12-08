<script setup>
import BmButton from '../components/atoms/BmButton.vue'
import BmInput from '../components/atoms/BmInput.vue'
import { ref } from 'vue'
import { useTokenStore } from '../store.js'
import { router } from '../router.js'

const email = ref('')
const password = ref('')
const showPassword = ref(false)


function login() {
  fetch('http://localhost:8000/api/login/', {
    method: "POST", // Or POST, PUT, DELETE, etc.
    body: JSON.stringify({
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
    const tokenStore = useTokenStore()
    tokenStore.setToken(data.data[0].token)
    router.push({ path: `/projects/${data.data[0].user.id}` });
  })
  .catch(error => {
    console.error('Error fetching data:', error);
  });
}

function signin() {
  router.push({ name: 'signin' });
}
  
  
</script>

<template>
<div class="font-display  bg-background-light  dark:bg-background-dark">
<div class="relative flex h-auto min-h-screen w-full flex-col bg-background-light dark:bg-background-dark group/design-root overflow-x-hidden">
<div class="layout-container flex h-full grow flex-col">
<div class="flex flex-1 justify-center items-stretch">
<div class="layout-content-container grid grid-cols-1 md:grid-cols-2 w-full">
<div class="relative hidden h-full w-full flex-col items-center justify-center bg-background-dark md:flex p-12 overflow-hidden">
<div class="absolute -top-16 -left-20 w-80 h-80 bg-primary/20 rounded-full blur-3xl"></div>
<div class="absolute -bottom-24 -right-24 w-96 h-96 bg-primary/20 rounded-full blur-3xl"></div>
<div class="absolute top-[20%] right-[10%] w-52 h-52 rounded-full border border-white/10 bg-white/5 backdrop-blur-md"></div>
<div class="relative z-10 flex flex-col items-center justify-center text-center">
<p class="text-4xl lg:text-5xl font-bold text-white leading-tight tracking-tight max-w-lg">Decodificando o Mundo de <span class="text-primary">Seis Patas</span>, Um Pixel de Cada Vez.</p>
</div>
</div>
<div class="relative flex w-full flex-col items-center justify-center bg-background-light dark:bg-background-dark p-6 sm:p-12 overflow-hidden">
<div class="absolute -bottom-1/4 -left-24 w-72 h-72 bg-primary/10 dark:bg-primary/5 rounded-full blur-3xl pointer-events-none"></div>
<div class="absolute -top-10 right-0 w-60 h-60 rounded-full border border-border-light dark:border-border-dark/50 bg-white/5 dark:bg-white/10 backdrop-blur-lg pointer-events-none -translate-x-1/4 translate-y-1/4"></div>
<div class="w-full max-w-md space-y-8 z-10">
<div class="text-center md:text-left">
<div class="flex items-center gap-3 justify-center md:justify-start mb-4">
<span class="material-symbols-outlined text-primary text-4xl">pest_control</span>
<h1 class="text-text-light dark:text-text-dark text-3xl font-bold">Bug Marker</h1>
</div>
<h2 class="text-text-light dark:text-text-dark tracking-light text-3xl font-bold leading-tight">Bem-vindo de Volta</h2>
<p class="text-subtle-text-light dark:text-subtle-text-dark mt-2">Faça login para continuar para o seu painel.</p>
</div>
<div class="flex flex-col gap-4">
<label class="flex flex-col w-full">
<p class="text-text-light dark:text-text-dark text-base font-medium leading-normal pb-2">Endereço de E-mail</p>
<BmInput
  type="email"
  placeholder="Digite seu e-mail"
  v-model="email"
/>
</label>
<label class="flex flex-col w-full">
<div class="flex justify-between items-baseline pb-2">
<p class="text-text-light dark:text-text-dark text-base font-medium leading-normal">Senha</p>
<a class="text-primary hover:text-opacity-80 dark:text-accent text-sm font-normal leading-normal underline hover:opacity-80" href="#">Esqueceu a senha?</a>
</div>
<div class="flex w-full flex-1 items-stretch rounded-lg">
<BmInput
  :type="showPassword ? 'text' : 'password'"
  placeholder="Digite sua senha"
  v-model="password"
  noRightBorder
/>
<div class="text-subtle-text-light dark:text-subtle-text-dark flex border border-border-light dark:border-border-dark bg-white dark:bg-background-dark items-center justify-center pr-[15px] rounded-r-lg border-l-0 cursor-pointer"
     @click="showPassword = !showPassword">
<span class="material-symbols-outlined">{{ showPassword ? 'visibility_off' : 'visibility' }}</span>
</div>
</div>
</label>
</div>
<BmButton @click="login" type="submit">Entrar</BmButton>
<p @click="signin"class="text-subtle-text-light dark:text-subtle-text-dark text-sm font-normal text-center">
                Não tem uma conta? <a class="font-bold text-primary dark:text-accent hover:underline" href="#">Cadastre-se</a>
</p>
</div>
</div>
</div>
</div>
</div>
</div>
</div>
</template>