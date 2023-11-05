<template>
  <div class="p-3">
    <Card class="overflow-auto main-chat-cont">
      <template #title>
        <div class="flex">
          <div class="col-12 mt-0">
            <p class="text-2xl my-0 py-0">Chat</p>
          </div>
        </div>
      </template>
      <template #content>
        <div v-if="hasStoredUsername">
          <Message v-for="msg in listOfMessages" style="width: 100%;" :key="msg.Message" severity="info" icon="pi pi-send"
            :closable="false">
            <div class="pl-2 flex flex-column">
              <div class="flex text-lg">{{ msg.Username + ': ' }}</div>
              <div class="flex align-left">{{ msg.Message }}</div>
            </div>
          </Message>
        </div>
      </template>
    </Card>


    <Toolbar class="mt-3 w-full">
      <template #start>
        <div class="mt-0">
          <p class="text-sm text-white mb-0 sm:mt-0">Chatting as:</p>
          <div v-if="hasStoredUsername">
            <p class="text-left mt-0 sm:mb-0"><span class="text-sm text-white font-semibold">{{ username }}</span></p>
          </div>
        </div>
      </template>
      <template #center>
        <div class="w-full w-auto-md">
          <InputText type="text" v-model="message" class="w-full" @keydown="sendOnEnter" />
        </div>
      </template>
      <template #end>
        <Button icon="pi pi-send" aria-label="Filter" @click="sendMessage" />
      </template>
    </Toolbar>
  </div>

  <Dialog v-if="!hasStoredUsername" class="dialog-width-responsive" v-model:visible="loginVisible" modal :pt="{
    mask: {
      style: 'backdrop-filter: blur(5px);'
    }
  }">
    <template #container="{ closeCallback }">
      <div class="flex flex-column px-7 py-7 gap-4"
        style="border-radius: 12px; background-image: radial-gradient(circle at left top, var(--primary-400), var(--primary-700))">
        <div class="flex justify-content-center">
          <img src="./assets/account.png" width="60" alt="">
        </div>
        <div class="inline-flex flex-column gap-2">
          <label for="username" class="text-primary-50 font-semibold">Username</label>
          <InputText id="username" v-model="username" class="bg-white-alpha-20 border-none p-3 text-primary-50" />
        </div>

        <div class="flex align-items-center gap-2">
          <Button label="Start Chatting" @click="loginCallback(closeCallback)" text
            class="p-3 w-full text-primary-50 border-1 border-white-alpha-30 hover:bg-white-alpha-10" />
        </div>
      </div>
    </template>
  </Dialog>
  <Toast />
</template>
<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue';
import { useToast } from "primevue/usetoast";
import ChatMsg from '@/model/Message'

// Use the environment variable if set, otherwise use "localhost"
const wsUrl = process.env.VUE_APP_WS_URL || "localhost";
const port = process.env.VUE_APP_PORT || "8080"
const socket = new WebSocket(`ws://${wsUrl}:${port}/ws`);

const message = ref("");
const listOfMessages = ref([] as ChatMsg[]);
const loginVisible = ref(true);
const username = ref("");
const toast = useToast();

const getLocalStorageValue = (key: string, defaultValue: any | null) => {
  const storedValue = localStorage.getItem(key);
  return storedValue ? storedValue : defaultValue;
};

const setLocalStorageValue = (key: string, value: any) => {
  localStorage.setItem(key, JSON.stringify(value));
};

const hasStoredUsername = ref(getLocalStorageValue('chat-username', null));

const checkLoginState = () => {
  const storedUsername = localStorage.getItem("chat-username");
  if (storedUsername) {
    username.value = storedUsername;
  }
}

onMounted(() => {
  checkLoginState();
});

const loginCallback = (closeCallback: () => void) => {
  if (!username.value) {
    toast.add({ severity: 'warn', summary: 'Missing Info', detail: 'Please provide a Username', life: 3000 });
    return;
  } else if (username.value.length > 10) {
    toast.add({ severity: 'warn', summary: 'Max-Length exceeded', detail: 'The length of the username provided is too long', life: 3000 });
    return;
  }
  localStorage.setItem("chat-username", username.value);
  hasStoredUsername.value = true;
  closeCallback();
}

const sendMessage = () => {
  if (message.value.length <= 0) return;
  const msg = new ChatMsg();
  msg.Username = username.value;
  msg.Message = message.value;

  socket.send(JSON.stringify(msg));
  message.value = "";
}

const sendOnEnter = (event: KeyboardEvent) => {
  if (event.key === "Enter") {
    sendMessage();
  }
}

socket.onopen = () => {
  console.log("Connection to WS-Server established");
}

socket.onmessage = (event) => {
  const msg: ChatMsg = JSON.parse(event.data);
  listOfMessages.value.push(msg);
}

socket.onclose = () => {
  console.log("Connection closed!");
}

socket.onerror = (err) => {
  console.error(err);
}

</script>


<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 30px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #42b983;
}

.main-chat-cont {
  height: 83vh;
}

.p-message-text {
  word-break: break-all;
}

.p-toolbar-group-center {
  width: 60%;
}

.dialog-width-responsive {
  width: 50vw;
}

@media (max-width: 768px) {
  .w-auto-md {
    width: auto;
  }

  .main-chat-cont {
    max-height: 69.5vh;
  }

  .p-toolbar-group-start {
    width: 100%;
    font-size: 1.3em !important;
  }

  .p-toolbar-group-center {
    width: 80%;
  }

  .p-toolbar-group-end {
    width: auto;
  }

  .dialog-width-responsive {
    width: 90vw;
  }
}

@media (max-width: 500px) {
  .main-chat-cont {
    max-height: 66vh;
  }
}
</style>
