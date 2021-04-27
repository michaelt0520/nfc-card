<template>
  <div
    v-if="isDisplay"
    class="relative flex flex-col sm:flex-row sm:items-center bg-white shadow rounded-md py-5 pl-6 pr-8 sm:pr-6"
  >
    <div
      class="flex flex-row items-center border-b sm:border-b-0 w-full sm:w-auto pb-4 sm:pb-0"
    >
      <div class="text-green-500">
        <ICOTick class="w-6 sm:w-5 h-6 sm:h-5" />
      </div>
      <div class="text-sm font-medium ml-3">Success.</div>
    </div>
    <div class="text-sm tracking-wide text-gray-500 mt-4 sm:mt-0 sm:ml-4">
      {{ message }}
    </div>
    <div
      class="absolute sm:relative sm:top-auto sm:right-auto ml-auto right-4 top-4 text-gray-400 hover:text-gray-800 cursor-pointer"
    >
      <ICODeny class="w-4 h-4" @click="$emit('update:isDisplay', false)" />
    </div>
  </div>
</template>

<script>
import ICOTick from "@/assets/icons/ICOTick";
import ICODeny from "@/assets/icons/ICODeny";

export default {
  name: "SuccessAlert",

  props: {
    message: {
      type: String,
      require: true,
      default: "Update success",
    },

    isDisplay: {
      type: Boolean,
      require: true,
      default: false,
    },
  },

  components: {
    ICOTick,
    ICODeny,
  },
};
</script>
