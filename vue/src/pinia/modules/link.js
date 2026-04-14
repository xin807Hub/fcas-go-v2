import { defineStore } from "pinia";
import { linkListApi } from "@/api/link.js";

const useLinkStore = defineStore("link", {
  state: () => ({
    linkData: [],
  }),
  // getters: {
  //   getLink: (state) => state.linkData,
  // },
  actions: {
    getLink() {
      return new Promise((resolve, reject) => {
        if (this.linkData.length > 0) {
          resolve(this.linkData);
        } else {
          linkListApi({ page: 1, limit: Number.MAX_SAFE_INTEGER, key: "" })
            .then((response) => {
              if (response.code == 0) {
                this.linkData = response.data.list;
              } else {
                this.linkData = [];
              }
              resolve(this.linkData);
            })
            .catch((err) => {
              console.log(err);
            });
        }
      });
    },
  },
});

export default useLinkStore;
