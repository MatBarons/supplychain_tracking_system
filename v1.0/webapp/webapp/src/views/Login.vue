<template>
    <main id="login">
        <div class="main-part">
            <div class="left-side-login">
            </div>
            <div class="right-side-login">
                <div class="response-login">
                    <h3>Loggati ora per gestire i tuoi lotti</h3>
                </div>
                <form @submit.prevent>
                    <label for="username">Email</label>
                    <input v-model="email" name="username" placeholder="Inserisci la tua email" required>
                    <label for="password">Password</label>
                    <input type="password" v-model="password" name="password" placeholder="Inserisci la tua password" required>
                    <button type="submit" class="btn-login" @click="logMeIn">Login</button>
                </form>
                <a href="url">Non hai un account? Registrati!</a>
                <div class="response-error">{{ response }}</div>
            </div>
        </div>
    </main>
</template>

<script lang="ts" setup>
import router from "@/router";
import { login } from "@/api_calls/users";
import { ref } from "vue";
import { useDataStore } from "@/stores/store";


const store = useDataStore()
const email = ref("")
const password = ref("")
const response = ref("")

function logMeIn(){
    login(email.value, password.value).then(response => {
        store.changeNomeAzienda(response.data.nomeAzienda)
        store.changePIVA(response.data.pIva)
        router.push('home')
    }).catch(reason => {
        response.value = "Credenziali errate"
    })
}
</script>

<style lang="scss" scoped>
#login {
    background: linear-gradient(90deg, #3b5998, #d1d8cc);
    overflow: hidden;
    height: 75.5%;
    padding-top: 15rem;
    padding-left: 23rem;
    .main-part {
        width: 60vw;
        height: 50vh;
        display: flex;
        box-shadow: 0 14px 28px rgba(0, 0, 0, 0.25), 0 10px 10px rgba(0, 0, 0, 0.22);
        border-radius: 1rem;
        overflow: hidden;

        .left-side-login {
            width: 40%;
            height: 100%;
            background-image: url("../assets/images/2_sfondo.jpg");
            background-size: cover;

            h1 {
                padding-top: 7rem;
                margin-left: 3rem;
                color: white;
                font-weight: bold;
                font-size: xx-large;
                padding-bottom: 0.5rem;
            }

            h3 {
                margin-left: 3rem;
                color: white;
                font-weight: bold;
            }

            img {
                margin-left: 15rem;
                margin-top: 5rem;
                width: 40%;
                height: 40%;
            }
        }

        .right-side-login {
            width: 60%;
            height: 100%;
            padding-left: 10rem;
            padding-top: 3rem;
            background: white;

            .response-login {
                padding-bottom: 1rem;
                //color: v-bind(colorResponse)
            }

            form {
                width: 100%;
                display: flex;
                flex-direction: column;

                label {
                    height: 15px;
                    padding-bottom: 0.5rem;
                }

                input {
                    font-size: 20px;
                    height: 3rem;
                    width: 20rem;
                    margin-bottom: 2rem;
                }

                .btn-login {
                    padding-top: 1rem;
                    padding-bottom: 1rem;
                    width: 30%;
                    color: white;
                    background-color: #00a817;
                    margin-top: 2rem;
                    margin-left: 5rem;
                    margin-bottom: 2rem;
                }
            }

            a {
                color: blue;
                margin-left: 3rem;
            }

            .response-error {
                color: red;
            }
        }
    }
}</style>