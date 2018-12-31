<template>
  <div id="app">
    <section class="tasks">
      <h1>
        Tasks 
        <transition name="fade">
          <small v-if="total">({{ total }})</small>
        </transition>
      </h1>
      <div class="tasks__new input-group">
        <input type="text" class="input-group-field" v-model="newTask" @keyup.enter="addTask" placeholder="New task">
        <span class="input-group-button">
          <button @click="addTask" class="button">
            <i>+</i>
          </button>
        </span>
      </div>

      <transition-group name="fade" tag="ul" class="tasks__list no-bullet">
        <ul class="tasks__list no-bullet" v-for="task in tasks" :key="JSON.stringify(task)">
          <li class="tasks__item" >
            <button
              class="tasks__item__toggle"
              @click.prevent="completeTask(task)"
              v-bind:class="{'tasks__item__toggle--done': task.done, 'tasks__item__toggle--completed': task.completed}"
            >✔️ {{ task.description }}</button>
          </li>
        </ul>
      </transition-group>
    </section>
  </div>
</template>

<script>
import axios from 'axios'
const baseURL = process.env.VUE_APP_API_URL
export default {
  name: 'app',
  data() {
    return {
      newTask: null,
      tasks: [],
      total: null
    }
  },

  async created() {
    await this.fetchData()
  },

  methods: {
    async addTask() {
      if (this.newTask && this.newTask.trim().length > 1) {
        const { data } = await axios.post(`${baseURL}/todos`, { description: this.newTask })
        this.tasks.unshift(data)
        this.newTask = null
        this.total++
      }
    },

    async fetchData() {
      const { data } = await axios.get(`${baseURL}/todos`)
      this.tasks = data
      const incompleted = data.filter(item => {
        return item.done ? null : item
      })
      this.total = incompleted.length
    },

    async completeTask(task) {
      task.completed = true
      const { data } = await axios.put(`${baseURL}/todos/${task.id}`)
      task.done = data.done
      this.total--
    }
  }
}
</script>

<style>
body {
  background-color: #abc;
}

*, h1, button {
  font-family: 'Nunito', sans-serif;
  font-size: 2rem;
}

h1 {
  margin: 0;
  padding: 0;
}
 
.fade-enter-active, .fade-leave-active {
  transition: opacity .5s;
}

.fade-enter, .fade-leave-active {
  opacity: 0;
}

input[type=text] {
  box-sizing: border-box;
  width: 100%;
  height: 2.4375rem;
  padding: .5rem;
  margin: 0 0 1rem;
  border: 1px solid #cacaca;
  font-size: 1.2rem;
  color: #0a0a0a;
  background-color: #fefefe;
  box-shadow: inset 0 1px 2px hsla(0,0%,4%,.1);
  border-radius: 0;
  -webkit-transition: -webkit-box-shadow .5s,border-color .25s ease-in-out;
  transition: box-shadow .5s,border-color .25s ease-in-out;
  -webkit-appearance: none;
  -moz-appearance: none;
  border-radius: 0 0 0 0;
  margin: 0;
  white-space: nowrap;
  display: table-cell;
  vertical-align: middle;
}

.input-group {
  display: table;
  width: 100%;
  margin-bottom: 1rem;
}

.input-group-field {
  border-radius: 0;
  height: 2.5rem;
}

input {
  line-height: normal;
}

.input-group .input-group-button {
  display: table-cell;
}

.input-group-button {
  padding-top: 0;
  padding-bottom: 0;
  text-align: center;
  height: 100%;
  width: 100px;
}
.input-group-button, .input-group-field, .input-group-label {
  margin: 0;
  white-space: nowrap;
  display: table-cell;
  vertical-align: middle;
}

.button {
  display: inline-block;
  text-align: center;
  line-height: 1;
  cursor: pointer;
  -webkit-appearance: none;
  -webkit-transition: background-color .25s ease-out,color .25s ease-out;
  transition: background-color .25s ease-out,color .25s ease-out;
  vertical-align: top;
  border: 1px solid transparent;
  border-radius: 0;
  padding: .2rem;
  width: 200px;
  font-size: .9rem;
  background-color: #2199e8;
  color: #fefefe;
}

.tasks__new {
  padding-top: 50px;
  padding-bottom: 20px;
}

.tasks {
  width: 100%;
  max-width: 45rem;
  padding: 1em;
  margin: 1em auto;
  overflow: auto;
  background-color: white;
  box-shadow: 0px .25rem 1rem rgba(black, .25);
}

ul {
  line-height: 1.6;
  margin-bottom: 1rem;
  list-style-position: outside;
  margin: 0;
  padding: 0;
}

.tasks__list {
  clear: both;
}

.tasks__item {
  margin-bottom: .5em;
  position: relative;
}

.tasks__item__toggle {
  cursor: pointer;
  width: 100%;
  text-align: left;
  padding: 0.85em 2.25em 0.85em 1em;
  background-color: rgba(0,0,0,0.05);
  border: 1px solid rgba(0,0,0,0.1);
  font-size: 1.2rem;
}

.tasks__item__toggle:hover {
  background-color: rgba(black, .1);
  border-color: rgba(black, .15);
}

.tasks__item__toggle--done {
  display: none;
}

.tasks__item__toggle--completed {
  text-decoration: line-through;
  background-color: rgba(green, .15);
  border-color: rgba(green, .2);
}

.no-bullet {
    margin-left: 0;
    list-style: none;
}
</style>
