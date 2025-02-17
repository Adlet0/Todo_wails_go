<template>
    <div>
      <input v-model="newTask" placeholder="Add new task" />
      <button @click="addTask">Add Task</button>
      <TaskItem
        v-for="(task, index) in tasks"
        :key="index"
        :task="task"
        :index="index"
        @remove="removeTask"
        @toggle="toggleTask"
      />
    </div>
  </template>
  
  <script>
  import TaskItem from './TaskItem.vue';
  
  export default {
    components: {
      TaskItem
    },
    data() {
      return {
        tasks: JSON.parse(localStorage.getItem('tasks')) || [],
        newTask: ''
      };
    },
    methods: {
      addTask() {
        if (this.newTask.trim() === '') return;
        const newTask = { Task: this.newTask, Done: false, Date: new Date() };
        this.tasks.push(newTask);
        this.newTask = '';
        this.saveTasks();
      },
      removeTask(index) {
        this.tasks.splice(index, 1);
        this.saveTasks();
      },
      toggleTask(index) {
        this.tasks[index].Done = !this.tasks[index].Done;
        this.saveTasks();
      },
      saveTasks() {
        localStorage.setItem('tasks', JSON.stringify(this.tasks));
      }
    }
  };
  </script>
  
  <style scoped>
  button {
    margin-left: 10px;
  }
  </style>
  