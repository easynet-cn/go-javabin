package javabin

const (
	/**
	 * Magic number that is written to the stream header.
	 */
	STREAM_MAGIC = 0xaced

	/**
	 * Version number that is written to the stream header.
	 */
	STREAM_VERSION = 5

	/* Each item in the stream is preceded by a tag
	 */

	/**
	 * First tag value.
	 */
	TC_BASE = 0x70

	/**
	 * Null object reference.
	 */
	TC_NULL = 0x70

	/**
	 * Reference to an object already written into the stream.
	 */
	TC_REFERENCE = 0x71

	/**
	 * new Class Descriptor.
	 */
	TC_CLASSDESC = 0x72

	/**
	 * new Object.
	 */
	TC_OBJECT = 0x73

	/**
	 * new String.
	 */
	TC_STRING = 0x74

	/**
	 * new Array.
	 */
	TC_ARRAY = 0x75

	/**
	 * Reference to Class.
	 */
	TC_CLASS = 0x76

	/**
	 * Block of optional data. Byte following tag indicates number
	 * of bytes in this block data.
	 */
	TC_BLOCKDATA = 0x77

	/**
	 * End of optional block data blocks for an object.
	 */
	TC_ENDBLOCKDATA = 0x78

	/**
	 * Reset stream context. All handles written into stream are reset.
	 */
	TC_RESET = 0x79

	/**
	 * long Block data. The long following the tag indicates the
	 * number of bytes in this block data.
	 */
	TC_BLOCKDATALONG = 0x7A

	/**
	 * Exception during write.
	 */
	TC_EXCEPTION = 0x7B

	/**
	 * Long string.
	 */
	TC_LONGSTRING = 0x7C

	/**
	 * new Proxy Class Descriptor.
	 */
	TC_PROXYCLASSDESC = 0x7D

	/**
	 * new Enum constant.
	 * @since 1.5
	 */
	TC_ENUM = 0x7E

	/**
	 * Last tag value.
	 */
	TC_MAX = 0x7E

	/******************************************************/
	/* Bit masks for ObjectStreamClass flag.*/

	/**
	 * Bit mask for ObjectStreamClass flag. Indicates a Serializable class
	 * defines its own writeObject method.
	 */
	SC_WRITE_METHOD = 0x01

	/**
	 * Bit mask for ObjectStreamClass flag. Indicates Externalizable data
	 * written in Block Data mode.
	 * Added for PROTOCOL_VERSION_2.
	 *
	 * @see #PROTOCOL_VERSION_2
	 * @since 1.2
	 */
	SC_BLOCK_DATA = 0x08

	/**
	 * Bit mask for ObjectStreamClass flag. Indicates class is Serializable.
	 */
	SC_SERIALIZABLE = 0x02

	/**
	 * Bit mask for ObjectStreamClass flag. Indicates class is Externalizable.
	 */
	SC_EXTERNALIZABLE = 0x04

	/**
	 * Bit mask for ObjectStreamClass flag. Indicates class is an enum type.
	 * @since 1.5
	 */
	SC_ENUM = 0x10

	/**
	 * A Stream Protocol Version. <p>
	 *
	 * All externalizable data is written in JDK 1.1 external data
	 * format after calling this method. This version is needed to write
	 * streams containing Externalizable data that can be read by
	 * pre-JDK 1.1.6 JVMs.
	 *
	 * @see java.io.ObjectOutputStream#useProtocolVersion(int)
	 * @since 1.2
	 */
	PROTOCOL_VERSION_1 = 1

	/**
	 * A Stream Protocol Version. <p>
	 *
	 * This protocol is written by JVM 1.2.
	 *
	 * Externalizable data is written in block data mode and is
	 * terminated with TC_ENDBLOCKDATA. Externalizable class descriptor
	 * flags has SC_BLOCK_DATA enabled. JVM 1.1.6 and greater can
	 * read this format change.
	 *
	 * Enables writing a nonSerializable class descriptor into the
	 * stream. The serialVersionUID of a nonSerializable class is
	 * set to 0L.
	 *
	 * @see java.io.ObjectOutputStream#useProtocolVersion(int)
	 * @see #SC_BLOCK_DATA
	 * @since 1.2
	 */
	PROTOCOL_VERSION_2 = 2
)
