<template>
  <div class="mx-auto container max-w-2xl shadow-md">
    <div
      class="bg-gray-100 p-4 border-t-2 bg-opacity-5 border-indigo-400 rounded-t"
    >
      <div class="max-w-sm mx-auto md:w-full md:mx-0">
        <div class="inline-flex items-center space-x-4">
          <h1 class="text-gray-600">Cards</h1>
        </div>
      </div>
    </div>

    <div class="bg-white space-y-6">
      <div class="flex pt-4 md:w-11/12 mx-auto">
        <input
          class="shadow appearance-none border border-gray-300 rounded-full w-full py-2 px-3 mx-4 h-10 text-grey-darker focus:outline-none"
          placeholder="Add Card"
          v-model="cardCode"
        />
        <button
          class="flex-no-shrink p-2 border-2 rounded text-teal border-teal hover:text-white hover:bg-teal"
          @click="onClickAddCard"
        >
          Add
        </button>
      </div>

      <div class="md:w-11/12 mx-auto py-5">
        <div
          class="flex mb-4 items-center"
          v-for="card in cardsList"
          :key="card.id"
        >
          <div class="w-full text-grey-darkest">
            <p class="flex flex-col">
              {{ card.code }}
            </p>
          </div>
          <label class="flex items-center">
            <input
              class="relative w-10 h-5 transition-all duration-200 ease-in-out bg-gray-400 rounded-full shadow-inner outline-none appearance-none cursor-pointer"
              type="checkbox"
              :checked="card.activated"
              @click="onClickToggleActivate(card)"
            />
          </label>
          <button
            class="flex-no-shrink p-2 ml-2 border-2 rounded text-red border-red hover:text-red-500 hover:bg-red"
            @click="onClickRemoveCard(card.code)"
          >
            <ICODelete />
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapActions } from "vuex";
import ICODelete from "@/assets/icons/ICODelete";

export default {
  name: "SettingCard",

  props: {
    cardsList: {
      type: Array,
      require: true,
    },
  },

  data() {
    return {
      cardCode: "",
    };
  },

  components: {
    ICODelete,
  },

  methods: {
    ...mapActions("users", ["addCard", "updateCard", "removeCard"]),

    onClickAddCard() {
      this.addCard({ code: this.cardCode });
    },

    onClickToggleActivate(card) {
      const data = {
        params: card.code,
        body: {
          activated: !card.activated,
        },
      };
      this.updateCard(data).then(() => {
        this.cardCode = "";
      });
    },

    onClickRemoveCard(cardCode) {
      this.removeCard(cardCode);
    },
  },
};
</script>

<style>
input:before {
  content: "";
  position: absolute;
  width: 1.25rem;
  height: 1.25rem;
  border-radius: 50%;
  top: 0;
  left: 0;
  transform: scale(1.1);
  box-shadow: 0 0.125rem 0.5rem rgba(0, 0, 0, 0.2);
  background-color: white;
  transition: 0.2s ease-in-out;
}

input:checked {
  background-color: #7f9cf5;
}

input:checked:before {
  left: 1.25rem;
}
</style>
