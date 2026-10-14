<template>
  <Head title="Targets" />
  <div class="targets-page">
    <header class="head">
      <h1 class="page-title">Targets</h1>
      <p class="lead">Add an endpoint to monitor and see everything you're watching.</p>
    </header>

    <form class="form card" @submit.prevent="handleSubmit">
      <h2 class="section-title">New target</h2>

      <div class="row">
        <label :class="{ invalid: form.errors.name }">
          <span class="label">Name</span>
          <input v-model="form.name" type="text" name="name" placeholder="Production API" />
          <span v-if="form.errors.name" class="error">{{ form.errors.name }}</span>
        </label>

        <label :class="{ invalid: form.errors.url }">
          <span class="label">URL</span>
          <input v-model="form.url" type="url" name="url" placeholder="https://api.example.com/health" />
          <span v-if="form.errors.url" class="error">{{ form.errors.url }}</span>
        </label>
      </div>

      <div class="row">
        <label :class="{ invalid: form.errors.interval }">
          <span class="label">Interval <span class="hint">seconds, min 10</span></span>
          <input v-model="form.interval" type="number" name="interval" min="10" />
          <span v-if="form.errors.interval" class="error">{{ form.errors.interval }}</span>
        </label>

        <label :class="{ invalid: form.errors.timeout }">
          <span class="label">Timeout <span class="hint">seconds, 1–30</span></span>
          <input v-model="form.timeout" type="number" name="timeout" min="1" max="30" />
          <span v-if="form.errors.timeout" class="error">{{ form.errors.timeout }}</span>
        </label>
      </div>

      <div class="actions">
        <button type="submit" class="btn" :disabled="form.processing">
          {{ form.processing ? 'Creating…' : 'Create target' }}
        </button>
      </div>
    </form>

    <section class="card list">
      <div class="list-head">
        <h2 class="section-title">Your targets</h2>
        <span class="count">{{ targets.length }}</span>
      </div>

      <p v-if="targets.length === 0" class="empty">No targets yet. Create one above to start monitoring.</p>

      <div v-else class="table-wrap">
        <table>
          <thead>
            <tr>
              <th>Name</th>
              <th>URL</th>
              <th class="num">Interval</th>
              <th class="num">Timeout</th>
              <th>Created</th>
              <th>Delete</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="target in targets" :key="target.id">
              <td class="name">{{ target.name }}</td>
              <td class="url">
                <a :href="target.url" target="_blank" rel="noopener noreferrer">{{ target.url }}</a>
              </td>
              <td class="num">{{ seconds(target.interval) }}</td>
              <td class="num">{{ seconds(target.timeout) }}</td>
              <td class="date">{{ formatDate(target.created_at) }}</td>
              <td>
                <Link preserve-scroll preserve-state class="text-red-600 cursor-pointer" :href="`targets/${target.id}/delete`" method="post" as="button">
                  <svg viewBox="0 0 24 24" width="24" height="24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M3 6h18"></path>
                    <path d="M19 6v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6"></path>
                    <path d="M8 6V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2"></path>
                    <line x1="10" y1="11" x2="10" y2="17"></line>
                    <line x1="14" y1="11" x2="14" y2="17"></line>
                  </svg>
                </Link>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
import { Head, useForm } from '@inertiajs/vue3'
import type { Target } from '../types'
import { Link } from '@inertiajs/vue3'

withDefaults(defineProps<{ targets?: Target[] }>(), {
  targets: () => [],
})

const form = useForm({
  name: '',
  url: '',
  interval: 30,
  timeout: 5,
})

const handleSubmit = () => {
  form.post('/targets', {
    preserveScroll: true,
    onSuccess: () => form.reset(),
  })
}

// Durations are serialized from Go's time.Duration, i.e. nanoseconds.
const seconds = (ns: number) => `${Math.round(ns / 1e9)}s`

const formatDate = (value: string) =>
  value ? new Date(value).toLocaleString(undefined, { dateStyle: 'medium', timeStyle: 'short' }) : '—'
</script>

<style scoped>
.targets-page {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}
.section-title {
  font-size: 1.1rem;
  font-weight: 700;
  letter-spacing: -0.01em;
}

/* form */
.form {
  display: flex;
  flex-direction: column;
  gap: 1.1rem;
}
.row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1rem;
}
label {
  display: flex;
  flex-direction: column;
  gap: 0.4rem;
}
.label {
  font-size: 0.85rem;
  font-weight: 600;
  color: var(--ink);
}
.hint {
  font-weight: 400;
  color: var(--muted);
  margin-left: 0.25rem;
}
input {
  font: inherit;
  padding: 0.65rem 0.8rem;
  border: 1px solid var(--border);
  border-radius: 10px;
  background: #fff;
  transition: border-color 0.15s, box-shadow 0.15s;
}
input:focus {
  outline: none;
  border-color: var(--go);
  box-shadow: 0 0 0 3px var(--ring);
}
label.invalid input {
  border-color: #e57373;
}
label.invalid input:focus {
  box-shadow: 0 0 0 3px rgba(229, 115, 115, 0.25);
}
.error {
  color: #b02a2a;
  font-size: 0.8rem;
}
.actions {
  display: flex;
  justify-content: flex-end;
  padding-top: 1.1rem;
  border-top: 1px solid var(--border);
}

/* table */
.list {
  padding: 0;
  overflow: hidden;
}
.list-head {
  display: flex;
  align-items: center;
  gap: 0.6rem;
  padding: 1.25rem 1.5rem;
  border-bottom: 1px solid var(--border);
}
.count {
  font-size: 0.78rem;
  font-weight: 600;
  color: var(--muted);
  background: #eef1f7;
  border-radius: 999px;
  padding: 0.1rem 0.55rem;
}
.empty {
  padding: 2.5rem 1.5rem;
  text-align: center;
  color: var(--muted);
}
.table-wrap {
  overflow-x: auto;
}
table {
  width: 100%;
  border-collapse: collapse;
  font-size: 0.9rem;
}
th,
td {
  padding: 0.75rem 1.5rem;
  text-align: left;
  white-space: nowrap;
}
th {
  font-size: 0.75rem;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.04em;
  color: var(--muted);
  background: #fafbfd;
  border-bottom: 1px solid var(--border);
}
tbody tr + tr td {
  border-top: 1px solid var(--border);
}
tbody tr:hover {
  background: #f9fafc;
}
.num {
  text-align: right;
  font-variant-numeric: tabular-nums;
}
.name {
  font-weight: 600;
}
.url {
  max-width: 280px;
  overflow: hidden;
  text-overflow: ellipsis;
}
.url a {
  color: var(--go);
  text-decoration: none;
}
.url a:hover {
  text-decoration: underline;
}
.date {
  color: var(--muted);
}

@media (max-width: 600px) {
  .row {
    grid-template-columns: 1fr;
  }
  th,
  td {
    padding: 0.7rem 1rem;
  }
}
</style>
