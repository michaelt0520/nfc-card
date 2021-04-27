<template>
  <div class="border-b border-grey-light p-4 bg-grey-lighter py-8">
    <div class="bg-white mx-auto max-w-sm shadow-lg rounded-lg overflow-hidden">
      <div class="sm:flex sm:items-center px-2 py-4">
        <div class="flex-grow">
          <h3 class="font-normal px-2 py-3 leading-tight">Contacts</h3>
          <input
            type="text"
            placeholder="Search teams or members"
            class="my-2 w-full text-sm bg-grey-light text-grey-darkest rounded h-10 p-3 focus:outline-none"
          />
          <div class="w-full">
            <div
              class="flex cursor-pointer my-1 hover:bg-blue-lightest rounded"
            >
              <div class="w-8 h-10 text-center py-1">
                <p class="text-3xl p-0 text-green-dark">&bull;</p>
              </div>
              <div class="w-4/5 h-10 py-3 px-1">
                <p class="hover:text-blue-dark">Kevin Durant</p>
              </div>
              <div class="w-1/5 h-10 text-right p-3">
                <p class="text-sm text-grey-dark">Member</p>
              </div>
            </div>
            <div
              class="flex cursor-pointer my-1 hover:bg-blue-lightest rounded"
            >
              <div class="w-8 h-10 text-center py-1">
                <p class="text-3xl p-0 text-grey-dark">&bull;</p>
              </div>
              <div class="w-4/5 h-10 py-3 px-1">
                <p class="hover:text-blue-dark">James Harden</p>
              </div>
              <div class="w-1/5 h-10 text-right p-3">
                <p class="text-sm text-grey-dark">Team</p>
              </div>
            </div>
            <div
              class="flex cursor-pointer my-1 hover:bg-blue-lightest rounded"
            >
              <div class="w-8 h-10 text-center py-1">
                <p class="text-3xl p-0 text-green-dark">&bull;</p>
              </div>
              <div class="w-4/5 h-10 py-3 px-1">
                <p class="hover:text-blue-dark">Michael Jordan</p>
              </div>
              <div class="w-1/5 h-10 text-right p-3">
                <p class="text-sm text-grey-dark">Member</p>
              </div>
            </div>
            <div
              class="flex cursor-pointer my-1 hover:bg-blue-lightest rounded"
            >
              <div class="w-8 h-10 text-center py-1">
                <p class="text-3xl p-0 text-grey-dark">&bull;</p>
              </div>
              <div class="w-4/5 h-10 py-3 px-1">
                <p class="hover:text-blue-dark">Tim Duncan</p>
              </div>
              <div class="w-1/5 h-10 text-right p-3">
                <p class="text-sm text-grey-dark">Team</p>
              </div>
            </div>
            <div
              class="flex my-1 cursor-pointer hover:bg-blue-lightest rounded"
            >
              <div class="w-8 h-10 text-center py-1">
                <p class="text-3xl p-0 text-grey-dark">&bull;</p>
              </div>
              <div class="w-4/5 h-10 py-3 px-1">
                <p class="hover:text-blue-dark">Lebrr James</p>
              </div>
              <div class="w-1/5 h-10 text-right p-3">
                <p class="text-sm text-grey-dark">Member</p>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="sm:flex bg-grey-light sm:items-center px-2 py-4">
        <div class="flex-grow text-right">
          <button
            class="text-grey-darkest hover:text-grey-dark py-2 px-4 rounded"
          >
            Cancel
          </button>
          <button
            class="bg-blue hover:bg-blue-dark text-white py-2 px-4 rounded"
          >
            Invite
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "PersonalUsersList",
};
</script>

<style>
</style>
