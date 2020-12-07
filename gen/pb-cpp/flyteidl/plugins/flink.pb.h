// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: flyteidl/plugins/flink.proto

#ifndef PROTOBUF_INCLUDED_flyteidl_2fplugins_2fflink_2eproto
#define PROTOBUF_INCLUDED_flyteidl_2fplugins_2fflink_2eproto

#include <limits>
#include <string>

#include <google/protobuf/port_def.inc>
#if PROTOBUF_VERSION < 3007000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers. Please update
#error your headers.
#endif
#if 3007000 < PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers. Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/port_undef.inc>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_table_driven.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/inlined_string_field.h>
#include <google/protobuf/metadata.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/map.h>  // IWYU pragma: export
#include <google/protobuf/map_entry.h>
#include <google/protobuf/map_field_inl.h>
#include <google/protobuf/unknown_field_set.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_flyteidl_2fplugins_2fflink_2eproto

// Internal implementation detail -- do not use these members.
struct TableStruct_flyteidl_2fplugins_2fflink_2eproto {
  static const ::google::protobuf::internal::ParseTableField entries[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::AuxillaryParseTableField aux[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::ParseTable schema[2]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::FieldMetadata field_metadata[];
  static const ::google::protobuf::internal::SerializationTable serialization_table[];
  static const ::google::protobuf::uint32 offsets[];
};
void AddDescriptors_flyteidl_2fplugins_2fflink_2eproto();
namespace flyteidl {
namespace plugins {
class FlinkJob;
class FlinkJobDefaultTypeInternal;
extern FlinkJobDefaultTypeInternal _FlinkJob_default_instance_;
class FlinkJob_FlinkPropertiesEntry_DoNotUse;
class FlinkJob_FlinkPropertiesEntry_DoNotUseDefaultTypeInternal;
extern FlinkJob_FlinkPropertiesEntry_DoNotUseDefaultTypeInternal _FlinkJob_FlinkPropertiesEntry_DoNotUse_default_instance_;
}  // namespace plugins
}  // namespace flyteidl
namespace google {
namespace protobuf {
template<> ::flyteidl::plugins::FlinkJob* Arena::CreateMaybeMessage<::flyteidl::plugins::FlinkJob>(Arena*);
template<> ::flyteidl::plugins::FlinkJob_FlinkPropertiesEntry_DoNotUse* Arena::CreateMaybeMessage<::flyteidl::plugins::FlinkJob_FlinkPropertiesEntry_DoNotUse>(Arena*);
}  // namespace protobuf
}  // namespace google
namespace flyteidl {
namespace plugins {

// ===================================================================

class FlinkJob_FlinkPropertiesEntry_DoNotUse : public ::google::protobuf::internal::MapEntry<FlinkJob_FlinkPropertiesEntry_DoNotUse, 
    ::std::string, ::std::string,
    ::google::protobuf::internal::WireFormatLite::TYPE_STRING,
    ::google::protobuf::internal::WireFormatLite::TYPE_STRING,
    0 > {
public:
#if GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
static bool _ParseMap(const char* begin, const char* end, void* object, ::google::protobuf::internal::ParseContext* ctx);
#endif  // GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  typedef ::google::protobuf::internal::MapEntry<FlinkJob_FlinkPropertiesEntry_DoNotUse, 
    ::std::string, ::std::string,
    ::google::protobuf::internal::WireFormatLite::TYPE_STRING,
    ::google::protobuf::internal::WireFormatLite::TYPE_STRING,
    0 > SuperType;
  FlinkJob_FlinkPropertiesEntry_DoNotUse();
  FlinkJob_FlinkPropertiesEntry_DoNotUse(::google::protobuf::Arena* arena);
  void MergeFrom(const FlinkJob_FlinkPropertiesEntry_DoNotUse& other);
  static const FlinkJob_FlinkPropertiesEntry_DoNotUse* internal_default_instance() { return reinterpret_cast<const FlinkJob_FlinkPropertiesEntry_DoNotUse*>(&_FlinkJob_FlinkPropertiesEntry_DoNotUse_default_instance_); }
  void MergeFrom(const ::google::protobuf::Message& other) final;
  ::google::protobuf::Metadata GetMetadata() const;
};

// -------------------------------------------------------------------

class FlinkJob final :
    public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:flyteidl.plugins.FlinkJob) */ {
 public:
  FlinkJob();
  virtual ~FlinkJob();

  FlinkJob(const FlinkJob& from);

  inline FlinkJob& operator=(const FlinkJob& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  FlinkJob(FlinkJob&& from) noexcept
    : FlinkJob() {
    *this = ::std::move(from);
  }

  inline FlinkJob& operator=(FlinkJob&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor() {
    return default_instance().GetDescriptor();
  }
  static const FlinkJob& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const FlinkJob* internal_default_instance() {
    return reinterpret_cast<const FlinkJob*>(
               &_FlinkJob_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    1;

  void Swap(FlinkJob* other);
  friend void swap(FlinkJob& a, FlinkJob& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline FlinkJob* New() const final {
    return CreateMaybeMessage<FlinkJob>(nullptr);
  }

  FlinkJob* New(::google::protobuf::Arena* arena) const final {
    return CreateMaybeMessage<FlinkJob>(arena);
  }
  void CopyFrom(const ::google::protobuf::Message& from) final;
  void MergeFrom(const ::google::protobuf::Message& from) final;
  void CopyFrom(const FlinkJob& from);
  void MergeFrom(const FlinkJob& from);
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  #if GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  static const char* _InternalParse(const char* begin, const char* end, void* object, ::google::protobuf::internal::ParseContext* ctx);
  ::google::protobuf::internal::ParseFunc _ParseFunc() const final { return _InternalParse; }
  #else
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) final;
  #endif  // GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const final;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      ::google::protobuf::uint8* target) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(FlinkJob* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return nullptr;
  }
  inline void* MaybeArenaPtr() const {
    return nullptr;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------


  // accessors -------------------------------------------------------

  // repeated string args = 3;
  int args_size() const;
  void clear_args();
  static const int kArgsFieldNumber = 3;
  const ::std::string& args(int index) const;
  ::std::string* mutable_args(int index);
  void set_args(int index, const ::std::string& value);
  #if LANG_CXX11
  void set_args(int index, ::std::string&& value);
  #endif
  void set_args(int index, const char* value);
  void set_args(int index, const char* value, size_t size);
  ::std::string* add_args();
  void add_args(const ::std::string& value);
  #if LANG_CXX11
  void add_args(::std::string&& value);
  #endif
  void add_args(const char* value);
  void add_args(const char* value, size_t size);
  const ::google::protobuf::RepeatedPtrField<::std::string>& args() const;
  ::google::protobuf::RepeatedPtrField<::std::string>* mutable_args();

  // map<string, string> flinkProperties = 4;
  int flinkproperties_size() const;
  void clear_flinkproperties();
  static const int kFlinkPropertiesFieldNumber = 4;
  const ::google::protobuf::Map< ::std::string, ::std::string >&
      flinkproperties() const;
  ::google::protobuf::Map< ::std::string, ::std::string >*
      mutable_flinkproperties();

  // string jarFile = 1;
  void clear_jarfile();
  static const int kJarFileFieldNumber = 1;
  const ::std::string& jarfile() const;
  void set_jarfile(const ::std::string& value);
  #if LANG_CXX11
  void set_jarfile(::std::string&& value);
  #endif
  void set_jarfile(const char* value);
  void set_jarfile(const char* value, size_t size);
  ::std::string* mutable_jarfile();
  ::std::string* release_jarfile();
  void set_allocated_jarfile(::std::string* jarfile);

  // string mainClass = 2;
  void clear_mainclass();
  static const int kMainClassFieldNumber = 2;
  const ::std::string& mainclass() const;
  void set_mainclass(const ::std::string& value);
  #if LANG_CXX11
  void set_mainclass(::std::string&& value);
  #endif
  void set_mainclass(const char* value);
  void set_mainclass(const char* value, size_t size);
  ::std::string* mutable_mainclass();
  ::std::string* release_mainclass();
  void set_allocated_mainclass(::std::string* mainclass);

  // @@protoc_insertion_point(class_scope:flyteidl.plugins.FlinkJob)
 private:
  class HasBitSetters;

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::RepeatedPtrField<::std::string> args_;
  ::google::protobuf::internal::MapField<
      FlinkJob_FlinkPropertiesEntry_DoNotUse,
      ::std::string, ::std::string,
      ::google::protobuf::internal::WireFormatLite::TYPE_STRING,
      ::google::protobuf::internal::WireFormatLite::TYPE_STRING,
      0 > flinkproperties_;
  ::google::protobuf::internal::ArenaStringPtr jarfile_;
  ::google::protobuf::internal::ArenaStringPtr mainclass_;
  mutable ::google::protobuf::internal::CachedSize _cached_size_;
  friend struct ::TableStruct_flyteidl_2fplugins_2fflink_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// -------------------------------------------------------------------

// FlinkJob

// string jarFile = 1;
inline void FlinkJob::clear_jarfile() {
  jarfile_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& FlinkJob::jarfile() const {
  // @@protoc_insertion_point(field_get:flyteidl.plugins.FlinkJob.jarFile)
  return jarfile_.GetNoArena();
}
inline void FlinkJob::set_jarfile(const ::std::string& value) {
  
  jarfile_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:flyteidl.plugins.FlinkJob.jarFile)
}
#if LANG_CXX11
inline void FlinkJob::set_jarfile(::std::string&& value) {
  
  jarfile_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:flyteidl.plugins.FlinkJob.jarFile)
}
#endif
inline void FlinkJob::set_jarfile(const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  
  jarfile_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:flyteidl.plugins.FlinkJob.jarFile)
}
inline void FlinkJob::set_jarfile(const char* value, size_t size) {
  
  jarfile_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:flyteidl.plugins.FlinkJob.jarFile)
}
inline ::std::string* FlinkJob::mutable_jarfile() {
  
  // @@protoc_insertion_point(field_mutable:flyteidl.plugins.FlinkJob.jarFile)
  return jarfile_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* FlinkJob::release_jarfile() {
  // @@protoc_insertion_point(field_release:flyteidl.plugins.FlinkJob.jarFile)
  
  return jarfile_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void FlinkJob::set_allocated_jarfile(::std::string* jarfile) {
  if (jarfile != nullptr) {
    
  } else {
    
  }
  jarfile_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), jarfile);
  // @@protoc_insertion_point(field_set_allocated:flyteidl.plugins.FlinkJob.jarFile)
}

// string mainClass = 2;
inline void FlinkJob::clear_mainclass() {
  mainclass_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& FlinkJob::mainclass() const {
  // @@protoc_insertion_point(field_get:flyteidl.plugins.FlinkJob.mainClass)
  return mainclass_.GetNoArena();
}
inline void FlinkJob::set_mainclass(const ::std::string& value) {
  
  mainclass_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:flyteidl.plugins.FlinkJob.mainClass)
}
#if LANG_CXX11
inline void FlinkJob::set_mainclass(::std::string&& value) {
  
  mainclass_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:flyteidl.plugins.FlinkJob.mainClass)
}
#endif
inline void FlinkJob::set_mainclass(const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  
  mainclass_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:flyteidl.plugins.FlinkJob.mainClass)
}
inline void FlinkJob::set_mainclass(const char* value, size_t size) {
  
  mainclass_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:flyteidl.plugins.FlinkJob.mainClass)
}
inline ::std::string* FlinkJob::mutable_mainclass() {
  
  // @@protoc_insertion_point(field_mutable:flyteidl.plugins.FlinkJob.mainClass)
  return mainclass_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* FlinkJob::release_mainclass() {
  // @@protoc_insertion_point(field_release:flyteidl.plugins.FlinkJob.mainClass)
  
  return mainclass_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void FlinkJob::set_allocated_mainclass(::std::string* mainclass) {
  if (mainclass != nullptr) {
    
  } else {
    
  }
  mainclass_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), mainclass);
  // @@protoc_insertion_point(field_set_allocated:flyteidl.plugins.FlinkJob.mainClass)
}

// repeated string args = 3;
inline int FlinkJob::args_size() const {
  return args_.size();
}
inline void FlinkJob::clear_args() {
  args_.Clear();
}
inline const ::std::string& FlinkJob::args(int index) const {
  // @@protoc_insertion_point(field_get:flyteidl.plugins.FlinkJob.args)
  return args_.Get(index);
}
inline ::std::string* FlinkJob::mutable_args(int index) {
  // @@protoc_insertion_point(field_mutable:flyteidl.plugins.FlinkJob.args)
  return args_.Mutable(index);
}
inline void FlinkJob::set_args(int index, const ::std::string& value) {
  // @@protoc_insertion_point(field_set:flyteidl.plugins.FlinkJob.args)
  args_.Mutable(index)->assign(value);
}
#if LANG_CXX11
inline void FlinkJob::set_args(int index, ::std::string&& value) {
  // @@protoc_insertion_point(field_set:flyteidl.plugins.FlinkJob.args)
  args_.Mutable(index)->assign(std::move(value));
}
#endif
inline void FlinkJob::set_args(int index, const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  args_.Mutable(index)->assign(value);
  // @@protoc_insertion_point(field_set_char:flyteidl.plugins.FlinkJob.args)
}
inline void FlinkJob::set_args(int index, const char* value, size_t size) {
  args_.Mutable(index)->assign(
    reinterpret_cast<const char*>(value), size);
  // @@protoc_insertion_point(field_set_pointer:flyteidl.plugins.FlinkJob.args)
}
inline ::std::string* FlinkJob::add_args() {
  // @@protoc_insertion_point(field_add_mutable:flyteidl.plugins.FlinkJob.args)
  return args_.Add();
}
inline void FlinkJob::add_args(const ::std::string& value) {
  args_.Add()->assign(value);
  // @@protoc_insertion_point(field_add:flyteidl.plugins.FlinkJob.args)
}
#if LANG_CXX11
inline void FlinkJob::add_args(::std::string&& value) {
  args_.Add(std::move(value));
  // @@protoc_insertion_point(field_add:flyteidl.plugins.FlinkJob.args)
}
#endif
inline void FlinkJob::add_args(const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  args_.Add()->assign(value);
  // @@protoc_insertion_point(field_add_char:flyteidl.plugins.FlinkJob.args)
}
inline void FlinkJob::add_args(const char* value, size_t size) {
  args_.Add()->assign(reinterpret_cast<const char*>(value), size);
  // @@protoc_insertion_point(field_add_pointer:flyteidl.plugins.FlinkJob.args)
}
inline const ::google::protobuf::RepeatedPtrField<::std::string>&
FlinkJob::args() const {
  // @@protoc_insertion_point(field_list:flyteidl.plugins.FlinkJob.args)
  return args_;
}
inline ::google::protobuf::RepeatedPtrField<::std::string>*
FlinkJob::mutable_args() {
  // @@protoc_insertion_point(field_mutable_list:flyteidl.plugins.FlinkJob.args)
  return &args_;
}

// map<string, string> flinkProperties = 4;
inline int FlinkJob::flinkproperties_size() const {
  return flinkproperties_.size();
}
inline void FlinkJob::clear_flinkproperties() {
  flinkproperties_.Clear();
}
inline const ::google::protobuf::Map< ::std::string, ::std::string >&
FlinkJob::flinkproperties() const {
  // @@protoc_insertion_point(field_map:flyteidl.plugins.FlinkJob.flinkProperties)
  return flinkproperties_.GetMap();
}
inline ::google::protobuf::Map< ::std::string, ::std::string >*
FlinkJob::mutable_flinkproperties() {
  // @@protoc_insertion_point(field_mutable_map:flyteidl.plugins.FlinkJob.flinkProperties)
  return flinkproperties_.MutableMap();
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__
// -------------------------------------------------------------------


// @@protoc_insertion_point(namespace_scope)

}  // namespace plugins
}  // namespace flyteidl

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // PROTOBUF_INCLUDED_flyteidl_2fplugins_2fflink_2eproto
