<template>
  <div class="antialiased font-sans bg-gray-200">
    <div class="container mx-auto px-4 sm:px-8">
      <div class="py-8">
        <div>
          <h2 class="text-2xl font-semibold leading-tight">Companies</h2>
        </div>
        <div class="my-2 flex sm:flex-row flex-col justify-between">
          <div class="flex flex-row">
            <div class="flex flex-row mb-1 sm:mb-0">
              <div class="relative">
                <select
                  class="h-full rounded-l border block appearance-none w-full bg-white border-gray-400 text-gray-700 py-2 px-4 pr-8 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                  @input="$emit('select-companies-per-page', $event)"
                >
                  <option>10</option>
                  <option>20</option>
                  <option>50</option>
                </select>
                <div
                  class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700"
                >
                  <svg
                    class="fill-current h-4 w-4"
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 20 20"
                  >
                    <path
                      d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z"
                    />
                  </svg>
                </div>
              </div>
            </div>
            <div class="block relative">
              <span
                class="h-full absolute inset-y-0 left-0 flex items-center pl-2"
              >
                <svg
                  viewBox="0 0 24 24"
                  class="h-4 w-4 fill-current text-gray-500"
                >
                  <path
                    d="M10 4a6 6 0 100 12 6 6 0 000-12zm-8 6a8 8 0 1114.32 4.906l5.387 5.387a1 1 0 01-1.414 1.414l-5.387-5.387A8 8 0 012 10z"
                  ></path>
                </svg>
              </span>
              <input
                placeholder="Search"
                class="appearance-none rounded-r rounded-l sm:rounded-l-none border border-gray-400 border-b block pl-8 pr-6 py-2 w-full bg-white text-sm placeholder-gray-400 text-gray-700 focus:bg-white focus:placeholder-gray-600 focus:text-gray-700 focus:outline-none"
                @input="$emit('search-company-input', $event)"
              />
            </div>
          </div>
          <div v-if="isShowButtonAdd">
            <button
              class="text-sm bg-green-400 hover:bg-green-500 text-gray-800 font-semibold py-2 px-4 rounded"
              @click="$emit('update:isOpenModal', true)"
            >
              Add
            </button>
          </div>
        </div>
        <div class="-mx-4 sm:-mx-8 px-4 sm:px-8 py-4 overflow-x-auto">
          <div
            class="inline-block min-w-full shadow rounded-lg overflow-hidden"
          >
            <table class="min-w-full leading-normal">
              <thead>
                <tr>
                  <th
                    class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
                  >
                    Name
                  </th>
                  <th
                    class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
                  >
                    Address
                  </th>
                  <th
                    class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
                  >
                    Website
                  </th>
                  <th
                    class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
                  >
                    Action
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="company in companies" :key="company.index">
                  <td
                    class="px-5 py-5 border-b border-gray-200 bg-white text-sm"
                  >
                    <div class="flex items-center">
                      <div class="flex-shrink-0 w-10 h-10">
                        <img
                          class="w-full h-full rounded-full"
                          :src="company.logo"
                          alt=""
                        />
                      </div>
                      <div class="ml-3">
                        <p class="text-gray-900 whitespace-no-wrap">
                          {{ company.name }}
                        </p>
                      </div>
                    </div>
                  </td>
                  <td
                    class="px-5 py-5 border-b border-gray-200 bg-white text-sm"
                  >
                    <p class="text-gray-900 whitespace-no-wrap">
                      {{ company.address }}
                    </p>
                  </td>
                  <td
                    class="px-5 py-5 border-b border-gray-200 bg-white text-sm"
                  >
                    <p class="text-gray-900 whitespace-no-wrap">
                      {{ company.website }}
                    </p>
                  </td>

                  <td
                    class="px-5 py-5 border-b border-gray-200 bg-white text-sm"
                  >
                    <div class="flex">
                      <button
                        class="flex-no-shrink p-2 border-2 rounded text-red border-red hover:text-red-500 hover:bg-green"
                        @click="onShowEditCompanyModal(company)"
                      >
                        <ICOEdit />
                      </button>

                      <button
                        class="flex-no-shrink p-2 ml-2 border-2 rounded text-red border-red hover:text-red-500 hover:bg-red"
                        @click="$emit('on-click-remove-company', company)"
                      >
                        <ICODelete />
                      </button>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
            <div
              class="px-5 py-5 bg-white border-t flex flex-col xs:flex-row items-center xs:justify-between"
            >
              <span class="text-xs xs:text-sm text-gray-900">
                Showing 1 to 4 of 50 Entries
              </span>
              <div class="inline-flex mt-2 xs:mt-0">
                <button
                  class="text-sm bg-gray-300 hover:bg-gray-400 text-gray-800 font-semibold py-2 px-4 rounded-l"
                >
                  Prev
                </button>
                <button
                  class="text-sm bg-gray-300 hover:bg-gray-400 text-gray-800 font-semibold py-2 px-4 rounded-r"
                >
                  Next
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <modal
      v-if="$route.name === 'admin'"
      v-model:open="isOpenModalEditCompany"
      header="Edit company"
      @confirm-modal="$emit('on-confirm-edit-company', editCompany)"
    >
      <template #modal-body>
        <div class="flex flex-row items-center justify-between py-4">
          <label for="name">Name</label>
          <input
            type="text"
            placeholder="name"
            class="w-4/5 text-sm bg-gray-200 text-gray-800 rounded h-10 p-3 focus:outline-none"
            v-model="editCompany.name"
          />
        </div>
        <div class="flex flex-row items-center justify-between py-4">
          <label for="address">Address</label>
          <input
            type="text"
            placeholder="Address"
            class="w-4/5 text-sm bg-gray-200 text-gray-800 rounded h-10 p-3 focus:outline-none"
            v-model="editCompany.address"
          />
        </div>
        <div class="flex flex-row items-center justify-between py-4">
          <label for="website">Homepage Website</label>
          <input
            type="text"
            placeholder="Homepage website"
            class="w-4/5 text-sm bg-gray-200 text-gray-800 rounded h-10 p-3 focus:outline-none"
            v-model="editCompany.website"
          />
        </div>
      </template>
    </modal>
  </div>
</template>

<script>
import ICODelete from "@/assets/icons/ICODelete";
import ICOEdit from "@/assets/icons/ICOEdit";
import Modal from "@/components/Modal";

export default {
  name: "Company",

  props: {
    companies: {
      type: Array,
      require: true,
    },

    isShowButtonAdd: {
      type: Boolean,
      default: false,
    },
  },

  components: {
    ICODelete,
    ICOEdit,
    Modal,
  },

  data() {
    return {
      isOpenModalEditCompany: false,
      editCompany: {},
    };
  },

  methods: {
    onShowEditCompanyModal(company) {
      console.log(company);
      this.editCompany = { ...company };
      this.isOpenModalEditCompany = true;
    },
  },
};
</script>
