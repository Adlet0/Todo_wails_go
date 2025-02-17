<template>
    <li>
      <span :class="{ done: task.Done }">{{ task.Task }}</span>
      <button @click="toggleTask">Toggle</button>
      <button @click="removeTask">Remove</button>
    </li>
  </template>
  
  <script>
  export default {
    props: ['task', 'index'],
    methods: {
      removeTask() {
        this.$emit('remove', this.index);
      },
      toggleTask() {
        this.$emit('toggle', this.index);
      }
    }
  };
  </script>
  
  <style scoped>
  .done {
    text-decoration: line-through;
    color: green;
  }
  button {
    margin-left: 10px;
  }
  </style>
  