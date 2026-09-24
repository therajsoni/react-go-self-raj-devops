import { useEffect, useState } from "react";

const emptyUser = {
  name: "",
  username: "",
  email: "",
  phoneNo: "",
  joiningDate: "",
  dob: "",
  officeTimeIn: "",
  officeTimeOut: "",
  photoBase64: "",
};

function UserDialog({ open, onClose, onSubmit, editingUser }) {
  const [form, setForm] = useState(emptyUser);

  useEffect(() => {
    if (editingUser) {
      setForm({
        name: editingUser.name || "",
        username: editingUser.username || "",
        email: editingUser.email || "",
        phoneNo: editingUser.phoneNo || "",
        joiningDate: editingUser.joiningDate || "",
        dob: editingUser.dob || "",
        officeTimeIn: editingUser.officeTimeIn || "",
        officeTimeOut: editingUser.officeTimeOut || "",
        photoBase64: editingUser.photoBase64 || "",
      });
    } else {
      setForm(emptyUser);
    }
  }, [editingUser, open]);

  function handleChange(e) {
    const { name, value } = e.target;

    setForm((prev) => ({
      ...prev,
      [name]: value,
    }));
  }

  function handlePhoto(e) {
    const file = e.target.files[0];

    if (!file) return;

    const reader = new FileReader();

    reader.onload = () => {
      setForm((prev) => ({
        ...prev,
        photoBase64: reader.result,
      }));
    };

    reader.readAsDataURL(file);
  }

  function handleSubmit(e) {
    e.preventDefault();

    onSubmit(form);
  }

  if (!open) {
    return null;
  }

  return (
    <div className="dialog-overlay">
      <div className="dialog">
        <div className="dialog-header">
          <h2>{editingUser ? "Update User" : "Add User"}</h2>

          <button className="close-btn" onClick={onClose}>
            ×
          </button>
        </div>

        <form onSubmit={handleSubmit}>
          <div className="form-grid">
            <div className="form-group">
              <label>Name</label>

              <input
                name="name"
                value={form.name}
                onChange={handleChange}
                required
              />
            </div>

            <div className="form-group">
              <label>Username</label>

              <input
                name="username"
                value={form.username}
                onChange={handleChange}
                required
              />
            </div>

            <div className="form-group">
              <label>Email</label>

              <input
                type="email"
                name="email"
                value={form.email}
                onChange={handleChange}
                required
              />
            </div>

            <div className="form-group">
              <label>Phone No</label>

              <input
                name="phoneNo"
                value={form.phoneNo}
                onChange={handleChange}
                required
              />
            </div>

            <div className="form-group">
              <label>Joining Date</label>

              <input
                type="date"
                name="joiningDate"
                value={form.joiningDate}
                onChange={handleChange}
                required
              />
            </div>

            <div className="form-group">
              <label>Date of Birth</label>

              <input
                type="date"
                name="dob"
                value={form.dob}
                onChange={handleChange}
                required
              />
            </div>

            <div className="form-group">
              <label>Office Time In</label>

              <input
                type="time"
                name="officeTimeIn"
                value={form.officeTimeIn}
                onChange={handleChange}
                required
              />
            </div>

            <div className="form-group">
              <label>Office Time Out</label>

              <input
                type="time"
                name="officeTimeOut"
                value={form.officeTimeOut}
                onChange={handleChange}
                required
              />
            </div>
          </div>

          <div className="form-group">
            <label>Photo</label>

            <input type="file" accept="image/*" onChange={handlePhoto} />
          </div>

          {form.photoBase64 && (
            <img
              className="photo-preview"
              src={form.photoBase64}
              alt="Preview"
            />
          )}

          <div className="dialog-actions">
            <button type="button" className="cancel-btn" onClick={onClose}>
              Cancel
            </button>

            <button type="submit">
              {editingUser ? "Update User" : "Create User"}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}

export default UserDialog;
