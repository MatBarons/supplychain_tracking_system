<template>
    <aside class="primary">
        <div class="infos">
            <img src="" alt="image" />
            <div>
                <h5>Welcome back,</h5>
                <h3>{{store.data.nomeAzienda}}</h3>
            </div>
        </div>
        <div class="menu">
            <router-link to="/home" class="button-top">
                <span class="material-icons">home</span>
                <span class="text">Home</span>
            </router-link>

            <div v-if="1 === 1"> <!-- Sostituire con una condizione di login-->
                <router-link to="/profile" class="button-top">
                    <span class="material-icons">person</span>
                    <span class="text">You</span>
                </router-link>
            </div>
            <div v-else>
                <router-link to="/login" class="button-top">
                    <span class="material-icons">person</span>
                    <span class="text">Login</span>
                </router-link>
            </div>
            <div class="bottom-part">
                <router-link to="/home" class="button-bottom" > <!-- cambiare il link -->
                    <span class="material-icons">anchor</span>
                    <span class="text">Support</span>
                </router-link>
                <router-link to="/" class="button-bottom" @click="logMeOut"> 
                    <span class="material-icons">power_settings_new</span>
                    <span class="text">Logout</span>
                </router-link>
            </div>
        </div>
    </aside>
</template>

<script lang="ts" setup>
import {useDataStore} from "@/stores/store"
import {logout} from "@/api_calls/users"
import { walletDisconnectionSocket } from "@/api_calls/socket";

const store = useDataStore();

async function logMeOut() {
    store.deleteAllData()
    logout().then(async response => {
        await walletDisconnectionSocket().then(res => {
            
        })
    })
}
</script>

<style lang="scss" scoped>
aside {
    position: relative;
    display: flex;
    flex-direction: column;
    overflow: hidden;
    min-width: 12%;
    height: auto;
    background: #3b5998;
    padding: 1.5rem;
    left: 0;
    @media(max-width: 1450px){
        width: 7%;
    }
    .infos{
        display: flex;
        align-items: center;
        margin-top: 1.5rem;
        img{
            color: black;
            margin-right: 1.5rem;
        }
        h5{
            color: white;
        }
        h3{
            color: white;
        }
    }
    .menu {
        margin: 2rem -1.5rem;
        @media(max-width: 1024px){
            width: 10%;
            
        }
        .button-top {
            display: flex;
            align-items: center;
            text-decoration: none;
            padding-left: 0.5rem;
            height: 4rem;
            margin-top: 2rem;
            .material-icons {
                font-size: 2rem;
                color: white;
            }
            .text {
                color: white;
                font-size: 1rem;
                padding-left: 1rem;
            }
            &:hover {
                background-color: var(--dark-alt);
                .material-icons,.text {
                    color: white;
                }
            }
            &.router-link-exact-active {
                background-color: var(--dark-alt);
                border-right: 5px solid black;
                .material-icons,.text {
                    color: black;
                }
            }
        }
        .bottom-part{
            position: absolute;
            bottom: 3%;
            padding-left: 17.5%;
            .button-bottom{
                display: flex;
                align-items: center;
                text-decoration: none;
                padding-left: 20%;
                height: 2rem;
                margin-top: 2rem;
                .material-icons {
                font-size: 1.5rem;
                color: white;
                }

                .text {
                    color: white;
                    font-size: 1rem;
                    padding-left: 1rem;
                }
                &:hover {
                    background-color: var(--dark-alt);

                  .material-icons,
                  .text {
                      color: white;
                  }
                }
            }
        }
    }

    @media (max-width: 1024px) {
        position: absolute;
        z-index: 99;
    }
}</style>