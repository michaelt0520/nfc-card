<template>
  <div>
    <div
      id="alert-success"
      class="fixed top-16 hidden right-5 flex flex-col sm:flex-row sm:items-center bg-white shadow rounded-md py-5 pl-6 pr-8 sm:pr-6"
    >
      <div
        class="flex flex-row items-center w-full sm:w-auto sm:pb-0"
      >
        <div class="text-green-500">
          <svg
            class="w-6 sm:w-5 h-6 sm:h-5"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"
            ></path>
          </svg>
        </div>
        <div class="text-sm font-medium ml-3">Success</div>
      </div>
    </div>
    <div
      id="alert-error"
      class="fixed top-16 hidden right-5 flex flex-col sm:flex-row sm:items-center bg-white shadow rounded-md py-5 pl-6 pr-8 sm:pr-6"
    >
      <div
        class="flex flex-row items-center w-full sm:w-auto sm:pb-0"
      >
        <div class="text-red-500">
          <svg
            class="w-6 sm:w-5 h-6 sm:h-5"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M6 18L18 6M6 6l12 12"
            ></path>
          </svg>
        </div>
        <div class="text-sm font-medium ml-3">Error</div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Alert",
};
</script>
