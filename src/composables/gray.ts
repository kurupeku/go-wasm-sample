const prefix = "data:image/png;base64,";

export const toBase64 = (file: File): Promise<string> => {
  return new Promise((resolve) => {
    const reader = new FileReader();
    reader.onload = (e) => {
      resolve(e.target.result.toString());
    };
    reader.readAsDataURL(file);
  });
};

export const grayScale = (data: string): Promise<string> => {
  return new Promise((resolve, reject) => {
    const base64 = data.replace(prefix, "");
    const res = window.grayScale(base64);
    if (res) resolve(`${prefix}${res}`);

    reject("could not convert source");
  });
};
