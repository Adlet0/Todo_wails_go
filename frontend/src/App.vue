<template>
  <div id="app">
    <h1>To-Do App</h1>
    <TaskList />
  </div>
</template>

<script>
// Импортируем компоненты
import TaskList from './components/TaskList.vue';

export default {
  components: {
    TaskList
  }
};
</script>

<style scoped>
h1 {
  font-size: 24px;
  text-align: center;
}
</style>
