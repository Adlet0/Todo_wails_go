<template>
    <div v-if="show" class="modal">
      <div class="modal-content">
        <p>Are you sure you want to delete this task?</p>
        <button @click="confirmDelete">Yes</button>
        <button @click="cancelDelete">No</button>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        show: false,
        taskToDelete: null
      };
    },
    methods: {
      confirmDelete() {
        this.$emit('confirm-delete', this.taskToDelete);
        this.show = false;
      },
      cancelDelete() {
        this.show = false;
      },
      openModal(taskIndex) {
        this.taskToDelete = taskIndex;
        this.show = true;
      }
    }
  };
  </script>
  
  <style scoped>
  .modal {
    position: fixed;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    background: white;
    padding: 20px;
    border: 1px solid black;
  }
  </style>
  